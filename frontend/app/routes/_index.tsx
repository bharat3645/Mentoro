import type { MetaFunction } from "@remix-run/node";
import { useState } from "react";
import BuddyChat from "~/components/buddy/BuddyChat";
import QuestManager from "~/components/quests/QuestManager";
import AnalyticsDashboard from "~/components/analytics/AnalyticsDashboard";
import Login from "~/components/login";
import { Quest, QuestStatus, QuestType } from "../components/quests/types";

export const meta: MetaFunction = () => {
  return [
    { title: "Learning Buddy Platform - Gamified AI Learning" },
    { name: "description", content: "Your AI-powered learning companion with gamified experience, adaptive personalities, and personalized quests!" },
  ];
};

// Mock data for demonstration
const mockUser = {
  name: "Alex",
  level: 5,
  xp: 750,
  xpToNext: 1000,
  streak: 7,
  totalBadges: 12,
  mood: "focused" as const,
};

const mockQuests: Quest[] = [
  { 
    id: 1, 
    title: "Debug Detective", 
    description: "Find and fix 3 bugs in your JavaScript code",
    progress: 2, 
    total: 3, 
    xp: 150, 
    type: 'debug' as const,
    difficulty: 2,
    status: 'active' as const,
    timeEstimate: 45
  },
  { 
    id: 2, 
    title: "Focus Flow Master", 
    description: "Complete a 25-minute focused coding session",
    progress: 20, 
    total: 25, 
    xp: 75, 
    type: 'focus' as const,
    difficulty: 1,
    status: 'active' as const,
    timeEstimate: 25
  },
  { 
    id: 3, 
    title: "React Hooks Explorer", 
    description: "Learn and implement 5 different React hooks",
    progress: 0, 
    total: 5, 
    xp: 200, 
    type: 'learn' as const,
    difficulty: 3,
    status: 'active' as const,
    timeEstimate: 120
  },
  { 
    id: 4, 
    title: "Code Review Champion", 
    description: "Review and provide feedback on 3 pull requests",
    progress: 3, 
    total: 3, 
    xp: 100, 
    type: 'code' as const,
    difficulty: 2,
    status: 'completed' as const,
    timeEstimate: 60
  },
];

const mockAnalytics = {
  totalXP: 2750,
  currentLevel: 5,
  streak: 7,
  questsCompleted: 12,
  badgesEarned: 8,
  studyTime: 420, // 7 hours in minutes
  weeklyProgress: [
    { day: 'Mon', xp: 150, time: 60 },
    { day: 'Tue', xp: 200, time: 90 },
    { day: 'Wed', xp: 100, time: 45 },
    { day: 'Thu', xp: 250, time: 120 },
    { day: 'Fri', xp: 180, time: 75 },
    { day: 'Sat', xp: 120, time: 30 },
    { day: 'Sun', xp: 0, time: 0 },
  ],
  moodData: [
    { date: '2024-01-15', mood: 'focused', performance: 85 },
    { date: '2024-01-14', mood: 'motivated', performance: 92 },
    { date: '2024-01-13', mood: 'tired', performance: 65 },
    { date: '2024-01-12', mood: 'excited', performance: 88 },
    { date: '2024-01-11', mood: 'focused', performance: 90 },
    { date: '2024-01-10', mood: 'happy', performance: 78 },
    { date: '2024-01-09', mood: 'chill', performance: 72 },
  ],
  skillProgress: [
    { skill: 'JavaScript', level: 7, maxLevel: 10 },
    { skill: 'React', level: 5, maxLevel: 10 },
    { skill: 'Node.js', level: 4, maxLevel: 10 },
    { skill: 'Database Design', level: 3, maxLevel: 10 },
    { skill: 'Problem Solving', level: 6, maxLevel: 10 },
  ],
};

const mockBadges = [
  { id: 1, name: "Code Warrior", icon: "⚔️", earned: true },
  { id: 2, name: "Focus Master", icon: "��", earned: true },
  { id: 3, name: "Streak Champion", icon: "🔥", earned: false },
  { id: 4, name: "Bug Hunter", icon: "🐛", earned: true },
  { id: 5, name: "Learning Machine", icon: "🧠", earned: true },
  { id: 6, name: "Social Learner", icon: "🤝", earned: false },
];

export default function Index() {
  const [currentView, setCurrentView] = useState<'dashboard' | 'quests' | 'analytics' | 'buddy'>('dashboard');
  const [buddyPersonality, setBuddyPersonality] = useState<'mentor' | 'cheerleader' | 'chill' | 'focused'>('focused');
  const [quests, setQuests] = useState<Quest[]>(mockQuests);

  const handleQuestUpdate = (questId: number, progress: number) => {
    setQuests(prev => prev.map(quest => 
      quest.id === questId ? { ...quest, progress } : quest
    ));
  };

  const handleQuestComplete = (questId: number) => {
    setQuests(prev => prev.map(quest => 
      quest.id === questId ? { ...quest, status: 'completed' as const } : quest
    ));
    // Here you would also update user XP, show celebration, etc.
  };

  const BuddyAvatar = ({ mood }: { mood: string }) => {
    const moodEmojis = {
      happy: "😊",
      focused: "🤓",
      chill: "😎",
      mentor: "👨‍🏫",
    };
    
    return (
      <div className="buddy-avatar flex items-center justify-center text-2xl bg-white">
        {moodEmojis[mood as keyof typeof moodEmojis] || "🤖"}
      </div>
    );
  };

  const XPBar = ({ current, max }: { current: number; max: number }) => {
    const percentage = (current / max) * 100;
    
    return (
      <div className="xp-bar">
        <div 
          className="xp-progress" 
          style={{ width: `${percentage}%` }}
        />
      </div>
    );
  };

  const BadgeGrid = ({ badges }: { badges: typeof mockBadges }) => {
    return (
      <div className="grid grid-cols-3 gap-2">
        {badges.slice(0, 6).map((badge) => (
          <div 
            key={badge.id}
            className={`p-2 rounded-lg text-center transition-all duration-200 ${
              badge.earned 
                ? 'bg-primary-100 border-2 border-primary-300 shadow-sm' 
                : 'bg-gray-100 border-2 border-gray-200 opacity-50'
            }`}
          >
            <div className="text-lg mb-1">{badge.icon}</div>
            <div className="text-xs font-medium text-gray-700">{badge.name}</div>
          </div>
        ))}
      </div>
    );
  };

  const Navigation = () => {
    const navItems = [
      { id: 'dashboard', label: 'Dashboard', icon: '🏠' },
      { id: 'quests', label: 'Quests', icon: '🎯' },
      { id: 'analytics', label: 'Analytics', icon: '📊' },
      { id: 'buddy', label: 'AI Buddy', icon: '🤖' },
      { id: 'login', label: 'Login', icon: '🔑', link: '/login' },
    ];

    return (
      <nav className="bg-white shadow-sm border-b border-gray-200 mb-8">
        <div className="max-w-6xl mx-auto px-4">
          <div className="flex space-x-8">
            {navItems.map((item) => (
              item.link ? (
                <a
                  key={item.id}
                  href={item.link}
                  className="flex items-center gap-2 py-4 px-2 border-b-2 border-transparent text-gray-500 hover:text-primary-600 transition-colors"
                >
                  <span>{item.icon}</span>
                  <span className="font-medium">{item.label}</span>
                </a>
              ) : (
                <button
                  key={item.id}
                  onClick={() => setCurrentView(item.id as any)}
                  className={`flex items-center gap-2 py-4 px-2 border-b-2 transition-colors ${
                    currentView === item.id
                      ? 'border-primary-500 text-primary-600'
                      : 'border-transparent text-gray-500 hover:text-gray-700'
                  }`}
                >
                  <span>{item.icon}</span>
                  <span className="font-medium">{item.label}</span>
                </button>
              )
            ))}
          </div>
        </div>
      </nav>
    );
  };

  const DashboardView = () => (
    <div className="space-y-8">
      {/* Header */}
      <div className="card">
        <div className="flex items-center justify-between">
          <div className="flex items-center gap-4">
            <BuddyAvatar mood={buddyPersonality} />
            <div>
              <h1 className="text-2xl font-bold text-gray-800">
                Welcome back, {mockUser.name}! 👋
              </h1>
              <p className="text-gray-600">Ready to continue your learning journey?</p>
            </div>
          </div>
          <div className="text-right">
            <div className="flex items-center gap-2 mb-1">
              <span className="text-lg">🔥</span>
              <span className="font-bold text-orange-600">{mockUser.streak} day streak</span>
            </div>
            <div className="text-sm text-gray-500">Keep it up!</div>
          </div>
        </div>
      </div>

      {/* Stats Grid */}
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        {/* Level & XP */}
        <div className="card">
          <h2 className="text-lg font-semibold mb-4 text-gray-800">Level Progress</h2>
          <div className="text-center mb-4">
            <div className="text-3xl font-bold text-primary-600">Level {mockUser.level}</div>
            <div className="text-sm text-gray-600">{mockUser.xp}/{mockUser.xpToNext} XP</div>
          </div>
          <XPBar current={mockUser.xp} max={mockUser.xpToNext} />
        </div>

        {/* Quick Stats */}
        <div className="card">
          <h2 className="text-lg font-semibold mb-4 text-gray-800">Quick Stats</h2>
          <div className="space-y-3">
            <div className="flex justify-between">
              <span className="text-gray-600">Active Quests</span>
              <span className="font-bold">{quests.filter(q => q.status === 'active').length}</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600">Completed Today</span>
              <span className="font-bold text-green-600">2</span>
            </div>
            <div className="flex justify-between">
              <span className="text-gray-600">Badges Earned</span>
              <span className="font-bold">{mockBadges.filter(b => b.earned).length}/{mockBadges.length}</span>
            </div>
          </div>
        </div>

        {/* Recent Badges */}
        <div className="card">
          <h2 className="text-lg font-semibold mb-4 text-gray-800">Recent Badges</h2>
          <BadgeGrid badges={mockBadges} />
          <button className="btn-primary w-full mt-4 text-sm">
            View All Achievements
          </button>
        </div>
      </div>

      {/* Recent Activity */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-8">
        {/* Active Quests Preview */}
        <div className="card">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-xl font-semibold text-gray-800">Active Quests</h2>
            <button 
              onClick={() => setCurrentView('quests')}
              className="text-sm text-primary-600 hover:text-primary-700"
            >
              View All →
            </button>
          </div>
          <div className="space-y-3">
            {quests.filter(q => q.status === 'active').slice(0, 3).map((quest) => {
              const progress = (quest.progress / quest.total) * 100;
              return (
                <div key={quest.id} className="p-3 bg-gray-50 rounded-lg">
                  <div className="flex justify-between items-start mb-2">
                    <h3 className="font-medium text-gray-800">{quest.title}</h3>
                    <span className="text-xs bg-primary-100 text-primary-700 px-2 py-1 rounded">
                      {quest.xp} XP
                    </span>
                  </div>
                  <div className="flex items-center gap-2">
                    <div className="flex-1 bg-gray-200 rounded-full h-2">
                      <div 
                        className="bg-primary-500 h-2 rounded-full transition-all duration-300"
                        style={{ width: `${progress}%` }}
                      />
                    </div>
                    <span className="text-xs text-gray-600">
                      {quest.progress}/{quest.total}
                    </span>
                  </div>
                </div>
              );
            })}
          </div>
        </div>

        {/* AI Buddy Preview */}
        <div className="card">
          <div className="flex justify-between items-center mb-4">
            <h2 className="text-xl font-semibold text-gray-800">AI Buddy</h2>
            <button 
              onClick={() => setCurrentView('buddy')}
              className="text-sm text-primary-600 hover:text-primary-700"
            >
              Open Chat →
            </button>
          </div>
          <div className="bg-gray-50 rounded-lg p-4">
            <div className="flex items-start gap-3">
              <BuddyAvatar mood={buddyPersonality} />
              <div className="flex-1">
                <div className="bg-white rounded-lg p-3 shadow-sm">
                  <p className="text-gray-800 text-sm">
                    Great job maintaining your 7-day streak! I can see you're in focused mode today. 
                    Ready to tackle that debugging quest? I can help you break it down into smaller steps.
                  </p>
                </div>
                <div className="text-xs text-gray-500 mt-1">Just now</div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-purple-50">
      <Navigation />
      
      <div className="max-w-6xl mx-auto p-4 md:p-8">
        {currentView === 'dashboard' && <DashboardView />}
        {currentView === 'quests' && (
          <QuestManager 
            quests={quests}
            onQuestUpdate={handleQuestUpdate}
            onQuestComplete={handleQuestComplete}
          />
        )}
        {currentView === 'analytics' && (
          <AnalyticsDashboard data={mockAnalytics} />
        )}
        {currentView === 'buddy' && (
          <div className="max-w-4xl mx-auto">
            <BuddyChat 
              personality={buddyPersonality}
              onPersonalityChange={(p) => setBuddyPersonality(p as any)}
            />
          </div>
        )}
      </div>
    </div>
  );
}

