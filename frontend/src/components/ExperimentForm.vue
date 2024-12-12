<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="实验名称">
      <el-input v-model="form.name" />
    </el-form-item>
    
    <el-form-item label="目标主机/容器">
      <el-input v-model="form.target" />
    </el-form-item>
    
    <el-form-item label="故障类型">
      <el-select v-model="form.action">
        <el-option label="CPU 负载" value="cpu-load" />
        <el-option label="内存压力" value="mem-load" />
        <el-option label="网络延迟" value="network-delay" />
        <el-option label="磁盘 IO" value="disk-burn" />
      </el-select>
    </el-form-item>
    
    <el-form-item label="参数配置" v-if="form.action">
      <template v-if="form.action === 'cpu-load'">
        <el-form-item label="CPU 占用率">
          <el-input-number v-model="form.parameters.percent" :min="0" :max="100" />
        </el-form-item>
      </template>
      
      <template v-if="form.action === 'network-delay'">
        <el-form-item label="延迟时间(ms)">
          <el-input-number v-model="form.parameters.time" :min="0" />
        </el-form-item>
      </template>
    </el-form-item>
    
    <el-form-item>
      <el-button type="primary" @click="submit">提交</el-button>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  data() {
    return {
      form: {
        name: '',
        target: '',
        action: '',
        parameters: {}
      }
    }
  },
  methods: {
    submit() {
      this.$emit('submit', this.form)
    }
  },
  watch: {
    'form.action': {
      handler(newAction) {
        this.form.parameters = {}
      }
    }
  }
}
</script> 