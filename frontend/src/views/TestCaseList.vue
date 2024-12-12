<template>
  <div class="test-case-list">
    <div class="page-header">
      <h2>故障注入用例</h2>
      <el-button type="primary" @click="createTestCase">新建用例</el-button>
    </div>

    <div class="content-wrapper">
      <!-- 左侧分类树 -->
      <div class="category-tree">
        <el-tree
          :data="categories"
          :props="{ label: 'name' }"
          @node-click="handleCategorySelect"
        >
          <template #default="{ node }">
            <div class="category-node">
              <span>{{ node.label }}</span>
              <span class="count">({{ getCategoryCount(node.id) }})</span>
            </div>
          </template>
        </el-tree>
      </div>

      <!-- 右侧用例列表 -->
      <div class="test-case-content">
        <!-- 搜索和筛选 -->
        <div class="filter-bar">
          <el-input
            v-model="searchQuery"
            placeholder="搜索用例"
            prefix-icon="el-icon-search"
            @input="handleSearch"
          />
          <el-select v-model="filterTags" multiple placeholder="标签筛选">
            <el-option
              v-for="tag in allTags"
              :key="tag"
              :label="tag"
              :value="tag"
            />
          </el-select>
        </div>

        <!-- 用例列表 -->
        <el-table :data="filteredTestCases" style="width: 100%">
          <el-table-column prop="name" label="用例名称">
            <template #default="{ row }">
              <div class="case-name">
                <span>{{ row.name }}</span>
                <el-tag
                  v-for="tag in row.tags"
                  :key="tag"
                  size="small"
                  class="ml-2"
                >
                  {{ tag }}
                </el-tag>
              </div>
            </template>
          </el-table-column>
          <el-table-column prop="category" label="分类" />
          <el-table-column prop="usageCount" label="使用次数" width="100" />
          <el-table-column label="收藏" width="80">
            <template #default="{ row }">
              <el-button
                :type="row.isFavorite ? 'warning' : 'info'"
                :icon="row.isFavorite ? 'Star' : 'StarFilled'"
                circle
                @click="toggleFavorite(row)"
              />
            </template>
          </el-table-column>
          <el-table-column label="操作" width="200">
            <template #default="{ row }">
              <el-button @click="runTestCase(row)">执行</el-button>
              <el-button @click="editTestCase(row)">编辑</el-button>
              <el-dropdown>
                <el-button icon="el-icon-more" />
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item @click="duplicateTestCase(row)">
                      复制
                    </el-dropdown-item>
                    <el-dropdown-item @click="deleteTestCase(row)">
                      删除
                    </el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </template>
          </el-table-column>
        </el-table>
      </div>
    </div>

    <!-- 用例编辑对话框 -->
    <el-dialog
      v-model="dialogVisible"
      :title="isEdit ? '编辑用例' : '新建用例'"
      width="60%"
    >
      <test-case-form
        :initial-data="currentTestCase"
        @submit="handleTestCaseSubmit"
        @cancel="dialogVisible = false"
      />
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue';
import { useStore } from 'vuex';
import TestCaseForm from '@/components/testcase/TestCaseForm.vue';

export default defineComponent({
  components: {
    TestCaseForm,
  },

  setup() {
    const store = useStore();
    const searchQuery = ref('');
    const filterTags = ref([]);
    const dialogVisible = ref(false);
    const isEdit = ref(false);
    const currentTestCase = ref(null);

    // 计算过滤后的用例列表
    const filteredTestCases = computed(() => {
      return store.getters['testCase/filteredTestCases']({
        search: searchQuery.value,
        tags: filterTags.value,
      });
    });

    // 处理收藏切换
    const toggleFavorite = async (testCase) => {
      await store.dispatch('testCase/toggleFavorite', testCase.id);
    };

    // 执行用例
    const runTestCase = (testCase) => {
      store.dispatch('experiment/createFromTestCase', testCase);
    };

    return {
      searchQuery,
      filterTags,
      dialogVisible,
      isEdit,
      currentTestCase,
      filteredTestCases,
      toggleFavorite,
      runTestCase,
    };
  },
});
</script>

<style lang="scss" scoped>
.test-case-list {
  padding: 20px;

  .content-wrapper {
    display: flex;
    margin-top: 20px;
    gap: 20px;
  }

  .category-tree {
    width: 250px;
    border-right: 1px solid #eee;
  }

  .test-case-content {
    flex: 1;
  }

  .filter-bar {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
  }

  .case-name {
    display: flex;
    align-items: center;
    gap: 8px;
  }
}
</style> 