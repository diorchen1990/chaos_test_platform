<template>
  <div class="report-detail">
    <div class="report-header">
      <h2>{{ report.name }}</h2>
      <div class="actions">
        <el-button @click="downloadReport('pdf')">下载 PDF</el-button>
        <el-button @click="downloadReport('html')">下载 HTML</el-button>
      </div>
    </div>

    <!-- 报告摘要 -->
    <el-card class="summary-card">
      <template #header>
        <div class="card-header">
          <span>实验摘要</span>
        </div>
      </template>
      <div class="summary-content">
        <div class="metric-item">
          <span class="label">总步骤数</span>
          <span class="value">{{ report.summary.totalSteps }}</span>
        </div>
        <div class="metric-item">
          <span class="label">成功步骤</span>
          <span class="value success">{{ report.summary.successSteps }}</span>
        </div>
        <div class="metric-item">
          <span class="label">失败步骤</span>
          <span class="value error">{{ report.summary.failedSteps }}</span>
        </div>
        <div class="metric-item">
          <span class="label">执行时长</span>
          <span class="value">{{ formatDuration(report.summary.duration) }}</span>
        </div>
      </div>
    </el-card>

    <!-- 执行步骤 -->
    <el-card class="steps-card">
      <template #header>
        <div class="card-header">
          <span>执行步骤</span>
        </div>
      </template>
      <el-timeline>
        <el-timeline-item
          v-for="step in report.steps"
          :key="step.name"
          :type="getStepType(step.status)"
          :timestamp="formatTime(step.startTime)"
        >
          <h4>{{ step.name }}</h4>
          <p>{{ step.action }} on {{ step.target }}</p>
          <p v-if="step.error" class="error-message">{{ step.error }}</p>
        </el-timeline-item>
      </el-timeline>
    </el-card>

    <!-- 监控指标 -->
    <el-card class="metrics-card">
      <template #header>
        <div class="card-header">
          <span>监控指标</span>
        </div>
      </template>
      <div class="metrics-charts">
        <div v-for="metric in report.metrics" :key="metric.name" class="chart-container">
          <h4>{{ metric.name }}</h4>
          <v-chart :option="getChartOption(metric)" />
        </div>
      </div>
    </el-card>

    <!-- 日志记录 -->
    <el-card class="logs-card">
      <template #header>
        <div class="card-header">
          <span>执行日志</span>
        </div>
      </template>
      <el-table :data="report.logs" height="400">
        <el-table-column prop="timestamp" label="时间" width="180" />
        <el-table-column prop="level" label="级别" width="100" />
        <el-table-column prop="source" label="来源" width="150" />
        <el-table-column prop="message" label="消息" />
      </el-table>
    </el-card>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { useStore } from 'vuex';
import VChart from 'vue-echarts';
import { formatDuration, formatTime } from '@/utils/format';

export default defineComponent({
  components: {
    VChart,
  },

  setup() {
    const route = useRoute();
    const store = useStore();
    const report = ref(null);

    onMounted(async () => {
      const reportId = route.params.id as string;
      report.value = await store.dispatch('report/getReport', reportId);
    });

    const downloadReport = async (format: string) => {
      await store.dispatch('report/downloadReport', {
        id: report.value.id,
        format,
      });
    };

    const getChartOption = (metric) => {
      return {
        title: {
          text: metric.name,
        },
        tooltip: {
          trigger: 'axis',
        },
        xAxis: {
          type: 'time',
        },
        yAxis: {
          type: 'value',
        },
        series: [{
          data: metric.values.map(v => [v.timestamp, v.value]),
          type: 'line',
        }],
      };
    };

    return {
      report,
      downloadReport,
      getChartOption,
      formatDuration,
      formatTime,
    };
  },
});
</script>

<style lang="scss" scoped>
.report-detail {
  padding: 20px;

  .report-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
  }

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .summary-content {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: 20px;
  }

  .metric-item {
    text-align: center;

    .label {
      font-size: 14px;
      color: #666;
    }

    .value {
      font-size: 24px;
      font-weight: bold;
      
      &.success {
        color: #67C23A;
      }
      
      &.error {
        color: #F56C6C;
      }
    }
  }

  .metrics-charts {
    display: grid;
    grid-template-columns: repeat(2, 1fr);
    gap: 20px;

    .chart-container {
      height: 300px;
    }
  }
}
</style> 