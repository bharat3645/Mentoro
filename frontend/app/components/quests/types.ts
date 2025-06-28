// Shared Quest types for Quests feature

export type QuestStatus = 'active' | 'completed' | 'paused' | 'failed';
export type QuestType = 'debug' | 'focus' | 'learn' | 'code';

export interface Quest {
  id: number;
  title: string;
  description: string;
  progress: number;
  total: number;
  xp: number;
  type: QuestType;
  difficulty: number;
  status: QuestStatus;
  timeEstimate: number;
} 