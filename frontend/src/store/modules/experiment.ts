import { Module } from 'vuex';
import { ExperimentState, RootState } from '../types';
import { Experiment } from '@/types';

const experimentModule: Module<ExperimentState, RootState> = {
    state: {
        experiments: [],
        loading: false,
        filters: {
            status: [],
            type: [],
            timeRange: null
        },
        pagination: {
            current: 1,
            pageSize: 10,
            total: 0
        }
    },
    
    getters: {
        filteredExperiments: (state) => {
            return state.experiments.filter(exp => {
                return true;
            });
        },
        
        paginatedExperiments: (state, getters) => {
            const { current, pageSize } = state.pagination;
            const filtered = getters.filteredExperiments;
            return filtered.slice((current - 1) * pageSize, current * pageSize);
        }
    }
}; 