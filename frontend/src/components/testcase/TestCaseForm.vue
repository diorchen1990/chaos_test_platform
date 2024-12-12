<template>
  <el-form :model="form" label-width="100px">
    <el-form-item label="用例名称" required>
      <el-input v-model="form.name" />
    </el-form-item>

    <el-form-item label="分类" required>
      <el-cascader
        v-model="form.category"
        :options="categories"
        :props="{ checkStrictly: true }"
      />
    </el-form-item>

    <el-form-item label="描述">
      <el-input
        v-model="form.description"
        type="textarea"
        :rows="3"
      />
    </el-form-item>

    <el-form-item label="标签">
      <el-select
        v-model="form.tags"
        multiple
        filterable
        allow-create
        default-first-option
        placeholder="请选择或创建标签"
      >
        <el-option
          v-for="tag in availableTags"
          :key="tag"
          :label="tag"
          :value="tag"
        />
      </el-select>
    </el-form-item>

    <el-form-item label="实验配置" required>
      <experiment-form v-model="form.experiment" />
    </el-form-item>

    <el-form-item>
      <el-button type="primary" @click="handleSubmit">保存</el-button>
      <el-button @click="$emit('cancel')">取消</el-button>
    </el-form-item>
  </el-form>
</template>

<script lang="ts">
import { defineComponent, ref, watch } from 'vue';
import ExperimentForm from '../experiment/ExperimentForm.vue';

export default defineComponent({
  components: {
    ExperimentForm,
  },

  props: {
    initialData: {
      type: Object,
      default: () => ({}),
    },
  },

  emits: ['submit', 'cancel'],

  setup(props, { emit }) {
    const form = ref({
      name: '',
      category: '',
      description: '',
      tags: [],
      experiment: null,
    });

    // 监听初始数据变化
    watch(
      () => props.initialData,
      (newVal) => {
        if (newVal) {
          form.value = { ...newVal };
        }
      },
      { immediate: true }
    );

    const handleSubmit = () => {
      emit('submit', form.value);
    };

    return {
      form,
      handleSubmit,
    };
  },
});
</script> 