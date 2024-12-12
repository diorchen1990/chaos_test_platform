<template>
  <div class="agent-management">
    <div class="page-header">
      <h2>探针管理</h2>
      <el-button type="primary" @click="showInstallDialog">
        安装探针
      </el-button>
    </div>

    <!-- 探针列表 -->
    <el-table :data="agents" v-loading="loading">
      <el-table-column prop="name" label="探针名称" />
      <el-table-column prop="type" label="类型" />
      <el-table-column prop="status" label="状态">
        <template #default="{ row }">
          <el-tag :type="getStatusType(row.status)">
            {{ row.status }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column prop="version" label="版本" />
      <el-table-column label="标签">
        <template #default="{ row }">
          <el-tag
            v-for="(value, key) in row.tags"
            :key="key"
            size="small"
            class="mx-1"
          >
            {{ key }}: {{ value }}
          </el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="250">
        <template #default="{ row }">
          <el-button 
            size="small"
            :type="row.status === 'running' ? 'danger' : 'success'"
            @click="handleAgentAction(row)"
          >
            {{ row.status === 'running' ? '停止' : '启动' }}
          </el-button>
          <el-button 
            size="small"
            type="primary"
            @click="handleUpdate(row)"
          >
            更新
          </el-button>
          <el-button 
            size="small"
            type="danger"
            @click="handleUninstall(row)"
          >
            卸载
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <!-- 安装对话框 -->
    <el-dialog
      v-model="installDialogVisible"
      title="安装探针"
      width="60%"
    >
      <agent-installer @success="handleInstallSuccess" />
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useStore } from 'vuex';
import AgentInstaller from '@/components/agent/AgentInstaller.vue';

export default defineComponent({
  name: 'AgentManagement',
  
  components: {
    AgentInstaller
  },

  setup() {
    const store = useStore();
    const loading = ref(false);
    const installDialogVisible = ref(false);

    const agents = computed(() => store.state.agent.agents);

    onMounted(async () => {
      loading.value = true;
      try {
        await store.dispatch('agent/fetchAgents');
      } finally {
        loading.value = false;
      }
    });

    const showInstallDialog = () => {
      installDialogVisible.value = true;
    };

    const handleInstallSuccess = () => {
      installDialogVisible.value = false;
      store.dispatch('agent/fetchAgents');
    };

    const handleAgentAction = async (agent) => {
      try {
        if (agent.status === 'running') {
          await store.dispatch('agent/stopAgent', agent.id);
        } else {
          await store.dispatch('agent/startAgent', agent.id);
        }
        await store.dispatch('agent/fetchAgents');
      } catch (error) {
        ElMessage.error('操作失败：' + error.message);
      }
    };

    const handleUpdate = async (agent) => {
      try {
        await store.dispatch('agent/updateAgent', agent.id);
        ElMessage.success('更新成功');
      } catch (error) {
        ElMessage.error('更新失败：' + error.message);
      }
    };

    const handleUninstall = async (agent) => {
      try {
        await ElMessageBox.confirm(
          '确定要卸载该探针吗？此操作不可恢复',
          '警告',
          {
            type: 'warning'
          }
        );
        await store.dispatch('agent/uninstallAgent', agent.id);
        ElMessage.success('卸载成功');
        await store.dispatch('agent/fetchAgents');
      } catch (error) {
        if (error !== 'cancel') {
          ElMessage.error('卸载失败：' + error.message);
        }
      }
    };

    const getStatusType = (status: string) => {
      const map = {
        running: 'success',
        stopped: 'info',
        error: 'danger'
      };
      return map[status] || 'info';
    };

    return {
      loading,
      agents,
      installDialogVisible,
      showInstallDialog,
      handleInstallSuccess,
      handleAgentAction,
      handleUpdate,
      handleUninstall,
      getStatusType
    };
  }
});
</script> 