import { useState } from "react";

interface AnalyticsData {
  totalXP: number;
  currentLevel: number;
  streak: number;
  questsCompleted: number;
  badgesEarned: number;
  studyTime: number; // in minutes
  weeklyProgress: Array<{ day: string; xp: number; time: number }>;
  moodData: Array<{ date: string; mood: string; performance: number }>;
  skillProgress: Array<{ skill: string; level: number; maxLevel: number }>;
}

interface AnalyticsDashboardProps {
  data: AnalyticsData;
}

const moodEmojis = {
  happy: "😊",
  focused: "🤓",
  excited: "🤩",
  tired: "😴",
  frustrated: "😤",
  confused: "😕",
  motivated: "💪",
  chill: "😎",
};

export default function AnalyticsDashboard({ data }: AnalyticsDashboardProps) {
  const [activeTab, setActiveTab] = useState<
    "overview" | "progress" | "mood" | "skills"
  >("overview");

  const ProgressBar = ({
    current,
    max,
    label,
    color = "bg-primary-500",
  }: {
    current: number;
    max: number;
    label: string;
    color?: string;
  }) => {
    const percentage = Math.min((current / max) * 100, 100);

    return (
      <div className="space-y-2">
        <div className="flex justify-between text-sm">
          <span className="text-gray-600">{label}</span>
          <span className="font-medium">
            {current}/{max}
          </span>
        </div>
        <div className="w-full bg-gray-200 rounded-full h-2">
          <div
            className={`h-2 rounded-full transition-all duration-500 ${color}`}
            style={{ width: `${percentage}%` }}
          />
        </div>
      </div>
    );
  };

  const StatCard = ({
    title,
    value,
    subtitle,
    icon,
    color = "text-primary-600",
  }: {
    title: string;
    value: string | number;
    subtitle?: string;
    icon: string;
    color?: string;
  }) => (
    <div className="card text-center">
      <div className="text-3xl mb-2">{icon}</div>
      <div className={`text-2xl font-bold ${color}`}>{value}</div>
      <div className="text-sm font-medium text-gray-700">{title}</div>
      {subtitle && <div className="text-xs text-gray-500 mt-1">{subtitle}</div>}
    </div>
  );

  const WeeklyChart = () => {
    const maxXP = Math.max(...data.weeklyProgress.map((d) => d.xp));

    return (
      <div className="card">
        <h3 className="text-lg font-semibold mb-4">Weekly Progress</h3>
        <div className="space-y-4">
          {data.weeklyProgress.map((day, index) => {
            const xpPercentage = maxXP > 0 ? (day.xp / maxXP) * 100 : 0;
            const timeHours = Math.round((day.time / 60) * 10) / 10;

            return (
              <div key={index} className="flex items-center gap-4">
                <div className="w-12 text-sm font-medium text-gray-600">
                  {day.day}
                </div>
                <div className="flex-1">
                  <div className="flex justify-between text-sm mb-1">
                    <span>{day.xp} XP</span>
                    <span>{timeHours}h</span>
                  </div>
                  <div className="w-full bg-gray-200 rounded-full h-2">
                    <div
                      className="h-2 bg-gradient-to-r from-primary-500 to-secondary-500 rounded-full transition-all duration-500"
                      style={{ width: `${xpPercentage}%` }}
                    />
                  </div>
                </div>
              </div>
            );
          })}
        </div>
      </div>
    );
  };

  const MoodTimeline = () => (
    <div className="card">
      <h3 className="text-lg font-semibold mb-4">
        Mood & Performance Timeline
      </h3>
      <div className="space-y-3">
        {data.moodData.slice(-7).map((entry, index) => (
          <div
            key={index}
            className="flex items-center justify-between p-3 bg-gray-50 rounded-lg"
          >
            <div className="flex items-center gap-3">
              <span className="text-2xl">
                {moodEmojis[entry.mood as keyof typeof moodEmojis] || "😐"}
              </span>
              <div>
                <div className="font-medium capitalize">{entry.mood}</div>
                <div className="text-sm text-gray-500">{entry.date}</div>
              </div>
            </div>
            <div className="text-right">
              <div className="font-semibold text-primary-600">
                {entry.performance}%
              </div>
              <div className="text-xs text-gray-500">Performance</div>
            </div>
          </div>
        ))}
      </div>
    </div>
  );

  const SkillsOverview = () => (
    <div className="card">
      <h3 className="text-lg font-semibold mb-4">Skill Progress</h3>
      <div className="space-y-4">
        {data.skillProgress.map((skill, index) => (
          <ProgressBar
            key={index}
            current={skill.level}
            max={skill.maxLevel}
            label={skill.skill}
            color={`bg-gradient-to-r from-primary-500 to-secondary-500`}
          />
        ))}
      </div>
    </div>
  );

  const OverviewTab = () => (
    <div className="space-y-6">
      {/* Key Stats */}
      <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
        <StatCard
          title="Total XP"
          value={data.totalXP.toLocaleString()}
          icon="⭐"
          color="text-yellow-600"
        />
        <StatCard
          title="Current Level"
          value={data.currentLevel}
          icon="📊"
          color="text-blue-600"
        />
        <StatCard
          title="Streak Days"
          value={data.streak}
          icon="🔥"
          color="text-orange-600"
        />
        <StatCard
          title="Study Time"
          value={`${Math.round(data.studyTime / 60)}h`}
          subtitle="This week"
          icon="⏱️"
          color="text-green-600"
        />
      </div>

      {/* Achievement Stats */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="card">
          <h3 className="text-lg font-semibold mb-4">Achievements</h3>
          <div className="space-y-3">
            <div className="flex justify-between items-center">
              <span className="text-gray-600">Quests Completed</span>
              <span className="font-bold text-primary-600">
                {data.questsCompleted}
              </span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-gray-600">Badges Earned</span>
              <span className="font-bold text-secondary-600">
                {data.badgesEarned}
              </span>
            </div>
            <div className="flex justify-between items-center">
              <span className="text-gray-600">Current Streak</span>
              <span className="font-bold text-orange-600">
                {data.streak} days
              </span>
            </div>
          </div>
        </div>

        <div className="card">
          <h3 className="text-lg font-semibold mb-4">Level Progress</h3>
          <div className="text-center mb-4">
            <div className="text-3xl font-bold text-primary-600">
              Level {data.currentLevel}
            </div>
            <div className="text-sm text-gray-500">Learning Champion</div>
          </div>
          <ProgressBar
            current={data.totalXP % 1000}
            max={1000}
            label="Progress to Next Level"
          />
        </div>
      </div>

      <WeeklyChart />
    </div>
  );

  const tabs = [
    { id: "overview", label: "Overview", icon: "📊" },
    { id: "progress", label: "Progress", icon: "📈" },
    { id: "mood", label: "Mood", icon: "😊" },
    { id: "skills", label: "Skills", icon: "🎯" },
  ];

  return (
    <div className="space-y-6">
      {/* Header */}
      <div className="flex justify-between items-center">
        <h2 className="text-2xl font-bold text-gray-800">
          Analytics Dashboard
        </h2>
        <div className="flex gap-2">
          <button className="text-sm bg-gray-100 text-gray-700 px-3 py-1 rounded-lg hover:bg-gray-200 transition-colors">
            Export Data
          </button>
          <button className="text-sm bg-primary-100 text-primary-700 px-3 py-1 rounded-lg hover:bg-primary-200 transition-colors">
            Share Progress
          </button>
        </div>
      </div>

      {/* Tab Navigation */}
      <div className="flex space-x-1 bg-gray-100 p-1 rounded-lg">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            onClick={() => setActiveTab(tab.id as any)}
            className={`flex-1 flex items-center justify-center gap-2 py-2 px-4 rounded-md transition-colors ${
              activeTab === tab.id
                ? "bg-white text-primary-600 shadow-sm"
                : "text-gray-600 hover:text-gray-800"
            }`}
          >
            <span>{tab.icon}</span>
            <span className="font-medium">{tab.label}</span>
          </button>
        ))}
      </div>

      {/* Tab Content */}
      <div className="min-h-96">
        {activeTab === "overview" && <OverviewTab />}
        {activeTab === "progress" && <WeeklyChart />}
        {activeTab === "mood" && <MoodTimeline />}
        {activeTab === "skills" && <SkillsOverview />}
      </div>
    </div>
  );
}
