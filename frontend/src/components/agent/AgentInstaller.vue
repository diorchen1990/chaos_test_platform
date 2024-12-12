<template>
  <div class="agent-installer">
    <el-steps :active="currentStep" finish-status="success">
      <el-step title="选择环境" />
      <el-step title="配置参数" />
      <el-step title="安装确认" />
    </el-steps>

    <!-- 步骤1：选择环境 -->
    <div v-if="currentStep === 0" class="step-content">
      <el-form :model="installConfig" label-width="120px">
        <el-form-item label="应用类型">
          <el-select v-model="installConfig.agentType">
            <el-option label="Java应用" value="java" />
            <el-option label="Go应用" value="golang" />
            <el-option label="Node.js应用" value="nodejs" />
          </el-select>
        </el-form-item>
        <el-form-item label="部署环境">
          <el-select v-model="installConfig.environment">
            <el-option label="物理机" value="host" />
            <el-option label="Docker" value="docker" />
            <el-option label="Kubernetes" value="kubernetes" />
          </el-select>
        </el-form-item>
      </el-form>
    </div>

    <!-- 步骤2：配置参数 -->
    <div v-if="currentStep === 1" class="step-content">
      <el-form :model="installConfig" label-width="120px">
        <el-form-item label="应用名称" required>
          <el-input v-model="installConfig.appName" />
        </el-form-item>
        <el-form-item label="服务地址" required>
          <el-input v-model="installConfig.serverUrl" />
        </el-form-item>
        <el-form-item label="标签">
          <el-tag
            v-for="tag in installConfig.tags"
            :key="tag.key"
            closable
            @close="removeTag(tag)"
          >
            {{ tag.key }}: {{ tag.value }}
          </el-tag>
          <el-button size="small" @click="showAddTag">添加标签</el-button>
        </el-form-item>
        <el-form-item label="监控配置">
          <el-switch
            v-model="installConfig.metrics.enable"
            active-text="启用"
            inactive-text="禁用"
          />
          <template v-if="installConfig.metrics.enable">
            <el-input-number
              v-model="installConfig.metrics.interval"
              :min="5"
              :max="60"
              label="采集间隔(秒)"
            />
          </template>
        </el-form-item>
      </el-form>
    </div>

    <!-- 步骤3：安装确认 -->
    <div v-if="currentStep === 2" class="step-content">
      <div class="install-command">
        <p>请在目标机器上执行以下命令：</p>
        <el-input
          type="textarea"
          :rows="4"
          v-model="installCommand"
          readonly
        />
        <el-button type="text" @click="copyCommand">
          复制命令
        </el-button>
      </div>
      <div class="config-preview">
        <p>配置预览：</p>
        <pre>{{ configYaml }}</pre>
      </div>
    </div>

    <!-- 步骤控制按钮 -->
    <div class="step-actions">
      <el-button @click="prevStep" v-if="currentStep > 0">上一步</el-button>
      <el-button 
        type="primary" 
        @click="nextStep"
        v-if="currentStep < 2"
      >
        下一步
      </el-button>
      <el-button 
        type="success" 
        @click="handleInstall"
        v-if="currentStep === 2"
      >
        开始安装
      </el-button>
    </div>

    <!-- 添加标签对话框 -->
    <el-dialog v-model="tagDialogVisible" title="添加标签">
      <el-form :model="newTag">
        <el-form-item label="键">
          <el-input v-model="newTag.key" />
        </el-form-item>
        <el-form-item label="值">
          <el-input v-model="newTag.value" />
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="tagDialogVisible = false">取消</el-button>
        <el-button type="primary" @click="addTag">确定</el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue';
import { useStore } from 'vuex';
import { ElMessage } from 'element-plus';
import yaml from 'js-yaml';

export default defineComponent({
  name: 'AgentInstaller',
  
  setup() {
    const store = useStore();
    const currentStep = ref(0);
    const tagDialogVisible = ref(false);
    
    const installConfig = ref({
      agentType: '',
      environment: '',
      appName: '',
      serverUrl: '',
      tags: [],
      metrics: {
        enable: true,
        interval: 15
      }
    });

    const newTag = ref({ key: '', value: '' });

    // 生成安装命令
    const installCommand = computed(() => {
      const baseUrl = installConfig.value.serverUrl;
      const type = installConfig.value.agentType;
      const env = installConfig.value.environment;
      
      return `curl -fsSL ${baseUrl}/install-agent.sh | bash -s -- \
        --agent-type=${type} \
        --app-name=${installConfig.value.appName} \
        --server-url=${baseUrl}`;
    });

    // 生成配置预览
    const configYaml = computed(() => {
      const config = {
        agent: {
          name: installConfig.value.appName,
          type: installConfig.value.agentType,
          server: installConfig.value.serverUrl,
          tags: Object.fromEntries(
            installConfig.value.tags.map(t => [t.key, t.value])
          )
        },
        metrics: installConfig.value.metrics
      };
      return yaml.dump(config);
    });

    // 复制命令
    const copyCommand = () => {
      navigator.clipboard.writeText(installCommand.value);
      ElMessage.success('命令已复制');
    };

    // 标签管理
    const showAddTag = () => {
      newTag.value = { key: '', value: '' };
      tagDialogVisible.value = true;
    };

    const addTag = () => {
      if (newTag.value.key && newTag.value.value) {
        installConfig.value.tags.push({ ...newTag.value });
        tagDialogVisible.value = false;
      }
    };

    const removeTag = (tag) => {
      const index = installConfig.value.tags.indexOf(tag);
      installConfig.value.tags.splice(index, 1);
    };

    // 步骤控制
    const nextStep = () => {
      if (validateCurrentStep()) {
        currentStep.value++;
      }
    };

    const prevStep = () => {
      currentStep.value--;
    };

    const validateCurrentStep = () => {
      if (currentStep.value === 1) {
        if (!installConfig.value.appName) {
          ElMessage.warning('请输入应用名称');
          return false;
        }
      }
      return true;
    };

    // 开始安装
    const handleInstall = async () => {
      try {
        await store.dispatch('agent/installAgent', installConfig.value);
        ElMessage.success('安装指令已生成，请在目标机器执行');
      } catch (error: any) {
        ElMessage.error(`安装失败：${error.message}`);
      }
    };

    return {
      currentStep,
      installConfig,
      tagDialogVisible,
      newTag,
      installCommand,
      configYaml,
      nextStep,
      prevStep,
      showAddTag,
      addTag,
      removeTag,
      copyCommand,
      handleInstall
    };
  }
});
</script>

<style lang="scss" scoped>
.agent-installer {
  padding: 20px;

  .step-content {
    margin: 20px 0;
    min-height: 300px;
  }

  .install-command {
    margin: 20px 0;
    
    pre {
      background: #f5f7fa;
      padding: 15px;
      border-radius: 4px;
    }
  }

  .config-preview {
    margin: 20px 0;
    
    pre {
      background: #f5f7fa;
      padding: 15px;
      border-radius: 4px;
    }
  }

  .step-actions {
    margin-top: 20px;
    text-align: right;
  }
}
</style> 