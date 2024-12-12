<template>
  <el-form :model="form" label-width="120px">
    <el-form-item label="Java进程名">
      <el-input v-model="form.target.javaApp.processName" />
    </el-form-item>
    
    <el-form-item label="故障类型">
      <el-select v-model="form.action">
        <el-option label="异常注入" value="java-exception" />
        <el-option label="GC压力" value="java-gc" />
        <el-option label="线程池满" value="java-threadpool" />
      </el-select>
    </el-form-item>
    
    <el-form-item label="故障参数" v-if="form.action">
      <template v-if="form.action === 'java-exception'">
        <el-input v-model="form.parameters.exceptionClass" placeholder="异常类名" />
        <el-input v-model="form.parameters.method" placeholder="目标方法" />
      </template>
      
      <template v-if="form.action === 'java-gc'">
        <el-input-number v-model="form.parameters.intervalMs" :min="1000" label="GC间隔(ms)" />
        <el-input-number v-model="form.parameters.duration" :min="1" label="持续时间(分钟)" />
      </template>
    </el-form-item>
  </el-form>
</template> 