export interface Agent {
    id: string;
    name: string;
    type: 'java' | 'golang' | 'nodejs';
    status: 'running' | 'stopped' | 'error';
    version: string;
    tags: Record<string, string>;
    metrics: {
        enable: boolean;
        interval: number;
    };
}

export interface AgentState {
    agents: Agent[];
    loading: boolean;
}

export interface AgentConfig {
    agentType: string;
    environment: string;
    appName: string;
    serverUrl: string;
    tags: Array<{ key: string; value: string }>;
    metrics: {
        enable: boolean;
        interval: number;
    };
} 