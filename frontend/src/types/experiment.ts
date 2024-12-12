export interface Experiment {
    id: string;
    name: string;
    scope: ExperimentScope;
    target: Target;
    action: string;
    parameters: any;
    status: string;
    createTime: string;
    updateTime: string;
    creator: string;
    agentIds: string[];
    results: ExperimentResult[];
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