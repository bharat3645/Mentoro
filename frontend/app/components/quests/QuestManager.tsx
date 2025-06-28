import { useState } from "react";
import { Quest, QuestStatus } from "./types";

interface QuestManagerProps {
  quests: Quest[];
  onQuestUpdate: (questId: number, progress: number) => void;
  onQuestComplete: (questId: number) => void;
}

const questTypeConfig = {
  code: { icon: '💻', color: 'bg-blue-100 text-blue-800', name: 'Coding' },
  focus: { icon: '🎯', color: 'bg-green-100 text-green-800', name: 'Focus' },
  learn: { icon: '📚', color: 'bg-purple-100 text-purple-800', name: 'Learning' },
  debug: { icon: '🐛', color: 'bg-red-100 text-red-800', name: 'Debug' },
  test: { icon: '🧪', color: 'bg-yellow-100 text-yellow-800', name: 'Testing' },
};

const difficultyStars = (difficulty: number) => {
  return '⭐'.repeat(difficulty) + '☆'.repeat(5 - difficulty);
};

export default function QuestManager({ quests, onQuestUpdate, onQuestComplete }: QuestManagerProps) {
  const [selectedQuest, setSelectedQuest] = useState<Quest | null>(null);
  const [showCompleted, setShowCompleted] = useState(false);

  const activeQuests = quests.filter(q => q.status === 'active');
  const completedQuests = quests.filter(q => q.status === 'completed');
  const displayQuests = showCompleted ? completedQuests : activeQuests;

  const handleProgressUpdate = (quest: Quest, increment: number) => {
    const newProgress = Math.min(quest.progress + increment, quest.total);
    onQuestUpdate(quest.id, newProgress);
    
    if (newProgress === quest.total) {
      onQuestComplete(quest.id);
    }
  };

  const QuestCard = ({ quest }: { quest: Quest }) => {
    const progressPercentage = (quest.progress / quest.total) * 100;
    const typeConfig = questTypeConfig[quest.type];
    const isCompleted = quest.status === 'completed';

    return (
      <div 
        className={`quest-item cursor-pointer transition-all duration-200 ${
          selectedQuest?.id === quest.id ? 'ring-2 ring-primary-500' : ''
        } ${isCompleted ? 'opacity-75' : ''}`}
        onClick={() => setSelectedQuest(quest)}
      >
        <div className="flex justify-between items-start mb-3">
          <div className="flex items-center gap-2">
            <span className="text-xl">{typeConfig.icon}</span>
            <div>
              <h3 className={`font-semibold ${isCompleted ? 'line-through text-gray-500' : 'text-gray-800'}`}>
                {quest.title}
              </h3>
              <p className="text-sm text-gray-600">{quest.description}</p>
            </div>
          </div>
          <div className="flex flex-col items-end gap-1">
            <span className={`badge ${typeConfig.color}`}>
              {quest.xp} XP
            </span>
            <div className="text-xs text-gray-500">
              {difficultyStars(quest.difficulty)}
            </div>
          </div>
        </div>

        <div className="space-y-2">
          {/* Progress Bar */}
          <div className="flex items-center gap-2">
            <div className="flex-1 bg-gray-200 rounded-full h-3">
              <div 
                className={`h-3 rounded-full transition-all duration-300 ${
                  isCompleted ? 'bg-green-500' : 'bg-primary-500'
                }`}
                style={{ width: `${progressPercentage}%` }}
              />
            </div>
            <span className="text-sm font-medium text-gray-600">
              {quest.progress}/{quest.total}
            </span>
          </div>

          {/* Quest Actions */}
          {!isCompleted && (
            <div className="flex justify-between items-center">
              <div className="flex gap-2">
                <button
                  onClick={(e) => {
                    e.stopPropagation();
                    handleProgressUpdate(quest, 1);
                  }}
                  className="text-xs bg-primary-100 text-primary-700 px-2 py-1 rounded hover:bg-primary-200 transition-colors"
                  disabled={quest.progress >= quest.total}
                >
                  +1 Progress
                </button>
                {quest.progress > 0 && (
                  <button
                    onClick={(e) => {
                      e.stopPropagation();
                      handleProgressUpdate(quest, -1);
                    }}
                    className="text-xs bg-gray-100 text-gray-700 px-2 py-1 rounded hover:bg-gray-200 transition-colors"
                  >
                    -1 Progress
                  </button>
                )}
              </div>
              
              {quest.timeEstimate && (
                <span className="text-xs text-gray-500 flex items-center gap-1">
                  ⏱️ {quest.timeEstimate}m
                </span>
              )}
            </div>
          )}

          {isCompleted && (
            <div className="flex items-center justify-center py-2">
              <span className="text-green-600 font-medium flex items-center gap-1">
                ✅ Completed! +{quest.xp} XP earned
              </span>
            </div>
          )}
        </div>
      </div>
    );
  };

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold text-gray-800">Quest Manager</h2>
        <div className="flex gap-2">
          <button
            onClick={() => setShowCompleted(false)}
            className={`px-4 py-2 rounded-lg transition-colors ${
              !showCompleted 
                ? 'bg-primary-500 text-white' 
                : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
            }`}
          >
            Active ({activeQuests.length})
          </button>
          <button
            onClick={() => setShowCompleted(true)}
            className={`px-4 py-2 rounded-lg transition-colors ${
              showCompleted 
                ? 'bg-primary-500 text-white' 
                : 'bg-gray-100 text-gray-700 hover:bg-gray-200'
            }`}
          >
            Completed ({completedQuests.length})
          </button>
        </div>
      </div>

      {/* Quest Stats */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="card text-center">
          <div className="text-2xl font-bold text-primary-600">{activeQuests.length}</div>
          <div className="text-sm text-gray-600">Active Quests</div>
        </div>
        <div className="card text-center">
          <div className="text-2xl font-bold text-green-600">{completedQuests.length}</div>
          <div className="text-sm text-gray-600">Completed</div>
        </div>
        <div className="card text-center">
          <div className="text-2xl font-bold text-orange-600">
            {quests.reduce((sum, q) => q.status === 'completed' ? sum + q.xp : sum, 0)}
          </div>
          <div className="text-sm text-gray-600">Total XP Earned</div>
        </div>
      </div>

      {/* Quest List */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {displayQuests.map((quest) => (
          <QuestCard key={quest.id} quest={quest} />
        ))}
      </div>

      {displayQuests.length === 0 && (
        <div className="text-center py-12">
          <div className="text-6xl mb-4">
            {showCompleted ? '🏆' : '🎯'}
          </div>
          <h3 className="text-xl font-semibold text-gray-700 mb-2">
            {showCompleted ? 'No completed quests yet' : 'No active quests'}
          </h3>
          <p className="text-gray-500">
            {showCompleted 
              ? 'Complete some quests to see them here!' 
              : 'Start a new quest to begin your learning journey!'
            }
          </p>
          {!showCompleted && (
            <button className="btn-primary mt-4">
              Generate New Quest
            </button>
          )}
        </div>
      )}

      {/* Quest Detail Modal */}
      {selectedQuest && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
          <div className="bg-white rounded-xl max-w-md w-full p-6">
            <div className="flex justify-between items-start mb-4">
              <h3 className="text-xl font-bold text-gray-800">{selectedQuest.title}</h3>
              <button
                onClick={() => setSelectedQuest(null)}
                className="text-gray-500 hover:text-gray-700"
              >
                ✕
              </button>
            </div>
            
            <div className="space-y-4">
              <p className="text-gray-600">{selectedQuest.description}</p>
              
              <div className="flex justify-between items-center">
                <span className="text-sm text-gray-500">Difficulty:</span>
                <span>{difficultyStars(selectedQuest.difficulty)}</span>
              </div>
              
              <div className="flex justify-between items-center">
                <span className="text-sm text-gray-500">XP Reward:</span>
                <span className="font-semibold text-primary-600">{selectedQuest.xp} XP</span>
              </div>
              
              <div className="flex justify-between items-center">
                <span className="text-sm text-gray-500">Progress:</span>
                <span className="font-semibold">{selectedQuest.progress}/{selectedQuest.total}</span>
              </div>
              
              {selectedQuest.timeEstimate && (
                <div className="flex justify-between items-center">
                  <span className="text-sm text-gray-500">Estimated Time:</span>
                  <span>{selectedQuest.timeEstimate} minutes</span>
                </div>
              )}
            </div>
            
            <div className="mt-6 flex gap-2">
              <button
                onClick={() => setSelectedQuest(null)}
                className="flex-1 bg-gray-100 text-gray-700 py-2 px-4 rounded-lg hover:bg-gray-200 transition-colors"
              >
                Close
              </button>
              {selectedQuest.status === 'active' && (
                <button
                  onClick={() => {
                    handleProgressUpdate(selectedQuest, 1);
                    setSelectedQuest(null);
                  }}
                  className="flex-1 btn-primary"
                >
                  Mark Progress
                </button>
              )}
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

