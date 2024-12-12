import { Module } from 'vuex';
import { ExperimentState, RootState } from '../types';
import { Experiment } from '@/types';

const experimentModule: Module<ExperimentState, RootState> = {
    state: {
        experiments: [],
        loading: false,
        currentExperiment: null,
        filters: {
            status: [],
            type: []
        }
    },
    
    getters: {
        filteredExperiments: (state) => {
            return state.experiments.filter(exp => {
                if (state.filters.status.length && 
                    !state.filters.status.includes(exp.status)) {
                    return false;
                }
                if (state.filters.type.length && 
                    !state.filters.type.includes(exp.type)) {
                    return false;
                }
                return true;
            });
        }
    }
}; 