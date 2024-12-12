import { Module } from 'vuex';
import { AgentState, RootState } from '../types';
import api from '@/api';

const agentModule: Module<AgentState, RootState> = {
  namespaced: true,

  state: {
    agents: [],
    loading: false
  },

  mutations: {
    SET_AGENTS(state, agents) {
      state.agents = agents;
    },
    SET_LOADING(state, loading) {
      state.loading = loading;
    }
  },

  actions: {
    async fetchAgents({ commit }) {
      commit('SET_LOADING', true);
      try {
        const response = await api.getAgents();
        commit('SET_AGENTS', response.data);
      } finally {
        commit('SET_LOADING', false);
      }
    },

    async installAgent({ dispatch }, config) {
      const response = await api.installAgent(config);
      return response.data;
    },

    async startAgent({ dispatch }, agentId) {
      await api.startAgent(agentId);
    },

    async stopAgent({ dispatch }, agentId) {
      await api.stopAgent(agentId);
    },

    async updateAgent({ dispatch }, agentId) {
      await api.updateAgent(agentId);
    },

    async uninstallAgent({ dispatch }, agentId) {
      await api.uninstallAgent(agentId);
    }
  }
};

export default agentModule; 