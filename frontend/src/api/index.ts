import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { Experiment, TestCase } from '@/types';

class API {
    private client: AxiosInstance;
    private baseURL: string;

    constructor() {
        this.baseURL = process.env.VUE_APP_API_URL || '/api';
        this.client = axios.create({
            baseURL: this.baseURL,
            timeout: 10000,
            headers: {
                'Content-Type': 'application/json',
            },
        });

        this.setupInterceptors();
    }

    private setupInterceptors() {
        this.client.interceptors.request.use(
            (config) => {
                const token = localStorage.getItem('token');
                if (token) {
                    config.headers.Authorization = `Bearer ${token}`;
                }
                return config;
            },
            (error) => {
                return Promise.reject(error);
            }
        );

        this.client.interceptors.response.use(
            (response) => response,
            (error) => {
                if (error.response?.status === 401) {
                    // 处理未授权
                    window.location.href = '/login';
                }
                return Promise.reject(error);
            }
        );
    }

    // 实验相关接口
    async getExperiments() {
        return this.client.get<Experiment[]>('/api/experiments');
    }

    async createExperiment(experiment: Partial<Experiment>) {
        return this.client.post<Experiment>('/api/experiments', experiment);
    }

    async getExperimentStatus(id: string) {
        return this.client.get(`/api/experiments/${id}/status`);
    }

    // 用例相关接口
    async getTestCases() {
        return this.client.get<TestCase[]>('/api/testcases');
    }

    async createTestCase(testCase: Partial<TestCase>) {
        return this.client.post<TestCase>('/api/testcases', testCase);
    }

    async toggleTestCaseFavorite(id: string) {
        return this.client.post(`/api/testcases/${id}/favorite`);
    }
}

export default new API(); 