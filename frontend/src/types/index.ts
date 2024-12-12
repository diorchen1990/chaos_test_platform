// 基础类型定义
export type ExperimentScope = 'host' | 'container' | 'kubernetes' | 'java-app' | 'database' | 'rpc';

export interface Experiment {
    id: string;
    name: string;
    scope: ExperimentScope;
    target: Target;
    action: string;
    parameters: Record<string, any>;
    status: string;
    createTime: string;
    updateTime: string;
    creator: string;
    results: ExperimentResult[];
}

export interface ExperimentResult {
    id: string;
    status: string;
    startTime: string;
    endTime: string;
    metrics: MetricData[];
    error?: string;
}

export interface TestCase {
    id: string;
    name: string;
    description: string;
    category: string;
    experiment: Experiment;
    favorite: boolean;
    createTime: string;
    updateTime: string;
}

export interface Report {
    id: string;
    experimentId: string;
    name: string;
    status: string;
    startTime: string;
    endTime: string;
    metrics: MetricData[];
    logs: LogEntry[];
}

export interface Target {
    type: string;
    endpoint?: string;
    labels?: Record<string, string>;
    javaApp?: JavaAppTarget;
    database?: DatabaseTarget;
    distributed?: DistributedTarget;
}

export interface JavaAppTarget {
    processName: string;
    port: number;
    classes?: string[];
    methods?: string[];
}

export interface DatabaseTarget {
    type: string;
    host: string;
    port: number;
    database: string;
    user: string;
    sqlPattern?: string;
}

export interface DistributedTarget {
    serviceName: string;
    instances: string[];
    apiPath?: string;
    dependencies?: string[];
}

// Store 状态类型定义
export interface RootState {
    version: string;
}

export interface ExperimentState {
    experiments: Experiment[];
    loading: boolean;
    currentExperiment: Experiment | null;
} 