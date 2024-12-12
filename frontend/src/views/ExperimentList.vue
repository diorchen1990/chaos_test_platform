<template>
  <div class="experiment-list">
    <el-button type="primary" @click="createExperiment">新建实验</el-button>
    
    <el-table :data="experiments">
      <el-table-column prop="name" label="实验名称" />
      <el-table-column prop="target" label="目标" />
      <el-table-column prop="action" label="故障类型" />
      <el-table-column prop="status" label="状态" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button @click="viewDetails(scope.row)">查看</el-button>
          <el-button type="danger" @click="stopExperiment(scope.row)">停止</el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <el-dialog v-model="dialogVisible" title="新建实验">
      <experiment-form @submit="handleSubmit" />
    </el-dialog>
  </div>
</template>

<script>
export default {
  data() {
    return {
      experiments: [],
      dialogVisible: false,
      timer: null
    }
  },
  methods: {
    async fetchExperiments() {
      const res = await this.$http.get('/api/experiments')
      this.experiments = res.data
    },
    createExperiment() {
      this.dialogVisible = true
    },
    async handleSubmit(data) {
      await this.$http.post('/api/experiments', data)
      this.dialogVisible = false
      this.fetchExperiments()
    },
    startStatusPolling() {
      this.timer = setInterval(() => {
        this.experiments.forEach(exp => {
          this.updateExperimentStatus(exp)
        })
      }, 5000)
    },
    async updateExperimentStatus(experiment) {
      try {
        const res = await this.$http.get(`/api/experiments/${experiment.id}/status`)
        const index = this.experiments.findIndex(e => e.id === experiment.id)
        if (index !== -1) {
          this.experiments[index].status = res.data.state
        }
      } catch (error) {
        console.error('更新状态失败:', error)
      }
    },
    async stopExperiment(experiment) {
      try {
        await this.$http.delete(`/api/experiments/${experiment.id}`)
        await this.fetchExperiments()
      } catch (error) {
        this.$message.error('停止实验失败')
      }
    }
  },
  mounted() {
    this.fetchExperiments()
    this.startStatusPolling()
  },
  beforeDestroy() {
    if (this.timer) {
      clearInterval(this.timer)
    }
  }
}
</script> 