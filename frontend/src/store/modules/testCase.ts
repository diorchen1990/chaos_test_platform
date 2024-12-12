import { Module } from 'vuex';
import { TestCaseState, RootState } from '../types';
import { TestCase } from '@/types';
import api from '@/api';

const testCaseModule: Module<TestCaseState, RootState> = {
  namespaced: true,

  state: {
    testCases: [],
    categories: [],
    loading: false,
  },

  getters: {
    filteredTestCases: (state) => (filters: { search: string; tags: string[] }) => {
      return state.testCases.filter((testCase) => {
        // 搜索过滤
        if (
          filters.search &&
          !testCase.name.toLowerCase().includes(filters.search.toLowerCase())
        ) {
          return false;
        }

        // 标签过滤
        if (
          filters.tags.length &&
          !filters.tags.some((tag) => testCase.tags.includes(tag))
        ) {
          return false;
        }

        return true;
      });
    },

    favoriteTestCases: (state) => {
      return state.testCases.filter((testCase) => testCase.isFavorite);
    },
  },

  mutations: {
    SET_TEST_CASES(state, testCases: TestCase[]) {
      state.testCases = testCases;
    },

    SET_CATEGORIES(state, categories) {
      state.categories = categories;
    },

    UPDATE_TEST_CASE(state, updatedTestCase: TestCase) {
      const index = state.testCases.findIndex((tc) => tc.id === updatedTestCase.id);
      if (index !== -1) {
        state.testCases[index] = updatedTestCase;
      }
    },
  },

  actions: {
    async fetchTestCases({ commit }) {
      const response = await api.getTestCases();
      commit('SET_TEST_CASES', response.data);
    },

    async toggleFavorite({ commit }, testCaseId: string) {
      const response = await api.toggleTestCaseFavorite(testCaseId);
      commit('UPDATE_TEST_CASE', response.data);
    },

    async createTestCase({ commit }, testCase: TestCase) {
      const response = await api.createTestCase(testCase);
      commit('SET_TEST_CASES', response.data);
    },
  },
};

export default testCaseModule; 