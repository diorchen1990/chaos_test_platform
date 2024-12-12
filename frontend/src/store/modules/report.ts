import { Module } from 'vuex';
import { ReportState, RootState } from '../types';
import api from '@/api';

const reportModule: Module<ReportState, RootState> = {
  namespaced: true,

  state: {
    reports: [],
    currentReport: null,
    loading: false,
  },

  mutations: {
    SET_REPORTS(state, reports) {
      state.reports = reports;
    },
    SET_CURRENT_REPORT(state, report) {
      state.currentReport = report;
    },
    SET_LOADING(state, loading) {
      state.loading = loading;
    },
  },

  actions: {
    async getReport({ commit }, id: string) {
      commit('SET_LOADING', true);
      try {
        const response = await api.getReport(id);
        commit('SET_CURRENT_REPORT', response.data);
        return response.data;
      } finally {
        commit('SET_LOADING', false);
      }
    },

    async downloadReport(_, { id, format }: { id: string; format: string }) {
      const response = await api.downloadReport(id, format);
      const blob = new Blob([response.data]);
      const url = window.URL.createObjectURL(blob);
      const link = document.createElement('a');
      link.href = url;
      link.download = `report_${id}.${format}`;
      link.click();
      window.URL.revokeObjectURL(url);
    },
  },
};

export default reportModule; 