import { useState, useRef, useEffect } from "react";

interface Message {
  id: string;
  type: 'user' | 'buddy';
  content: string;
  timestamp: Date;
  personality?: string;
}

interface BuddyChatProps {
  personality: 'mentor' | 'cheerleader' | 'chill' | 'focused';
  onPersonalityChange: (personality: string) => void;
}

const personalityConfig = {
  mentor: {
    name: "Wise Mentor",
    emoji: "👨‍🏫",
    color: "text-blue-600",
    bgColor: "bg-blue-50",
  },
  cheerleader: {
    name: "Enthusiastic Cheerleader", 
    emoji: "🎉",
    color: "text-pink-600",
    bgColor: "bg-pink-50",
  },
  chill: {
    name: "Chill Friend",
    emoji: "😎",
    color: "text-purple-600", 
    bgColor: "bg-purple-50",
  },
  focused: {
    name: "Focused Analyst",
    emoji: "🤓",
    color: "text-green-600",
    bgColor: "bg-green-50",
  },
};

export default function BuddyChat({ personality, onPersonalityChange }: BuddyChatProps) {
  const [messages, setMessages] = useState<Message[]>([
    {
      id: '1',
      type: 'buddy',
      content: "Hey there! I'm your AI learning buddy. How can I help you today?",
      timestamp: new Date(),
      personality: personality,
    }
  ]);
  const [inputValue, setInputValue] = useState('');
  const [isTyping, setIsTyping] = useState(false);
  const messagesEndRef = useRef<HTMLDivElement>(null);

  const scrollToBottom = () => {
    messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const handleSendMessage = async () => {
    if (!inputValue.trim()) return;

    const userMessage: Message = {
      id: Date.now().toString(),
      type: 'user',
      content: inputValue,
      timestamp: new Date(),
    };

    setMessages(prev => [...prev, userMessage]);
    setInputValue('');
    setIsTyping(true);

    // Simulate AI response
    setTimeout(() => {
      const responses = {
        mentor: [
          "That's a great question! Let me break this down step by step for you.",
          "I can see you're thinking deeply about this. Here's how I'd approach it...",
          "Excellent! You're making real progress. Let's build on that understanding.",
        ],
        cheerleader: [
          "You're doing amazing! 🎉 Keep up that fantastic energy!",
          "I love your enthusiasm! You're absolutely crushing this! 💪",
          "Way to go! Your progress is incredible and I'm so proud of you! ⭐",
        ],
        chill: [
          "No worries at all! Let's just take this one step at a time.",
          "Hey, that's totally cool. We'll figure this out together, no pressure.",
          "All good! Sometimes the best insights come when we're relaxed.",
        ],
        focused: [
          "Let's analyze this systematically. Here are the key factors to consider...",
          "Good observation. Let's dive deeper into the technical details.",
          "Precisely! Now let's optimize this approach for maximum efficiency.",
        ],
      };

      const responseList = responses[personality];
      const randomResponse = responseList[Math.floor(Math.random() * responseList.length)];

      const buddyMessage: Message = {
        id: (Date.now() + 1).toString(),
        type: 'buddy',
        content: randomResponse,
        timestamp: new Date(),
        personality: personality,
      };

      setMessages(prev => [...prev, buddyMessage]);
      setIsTyping(false);
    }, 1500);
  };

  const handleKeyPress = (e: React.KeyboardEvent) => {
    if (e.key === 'Enter' && !e.shiftKey) {
      e.preventDefault();
      handleSendMessage();
    }
  };

  const config = personalityConfig[personality];

  return (
    <div className="card h-96 flex flex-col">
      {/* Header */}
      <div className="flex items-center justify-between mb-4 pb-3 border-b border-gray-200">
        <div className="flex items-center gap-3">
          <div className={`w-10 h-10 rounded-full ${config.bgColor} flex items-center justify-center text-lg`}>
            {config.emoji}
          </div>
          <div>
            <h3 className="font-semibold text-gray-800">{config.name}</h3>
            <p className="text-sm text-gray-500">AI Learning Buddy</p>
          </div>
        </div>
        
        {/* Personality Switcher */}
        <select
          value={personality}
          onChange={(e) => onPersonalityChange(e.target.value)}
          className="text-sm border border-gray-300 rounded-md px-2 py-1 focus:outline-none focus:ring-2 focus:ring-primary-500"
        >
          <option value="mentor">👨‍🏫 Mentor</option>
          <option value="cheerleader">🎉 Cheerleader</option>
          <option value="chill">😎 Chill</option>
          <option value="focused">🤓 Focused</option>
        </select>
      </div>

      {/* Messages */}
      <div className="flex-1 overflow-y-auto space-y-3 mb-4">
        {messages.map((message) => (
          <div
            key={message.id}
            className={`flex ${message.type === 'user' ? 'justify-end' : 'justify-start'}`}
          >
            <div
              className={`max-w-xs lg:max-w-md px-4 py-2 rounded-lg ${
                message.type === 'user'
                  ? 'bg-primary-500 text-white'
                  : `${config.bgColor} ${config.color}`
              }`}
            >
              <p className="text-sm">{message.content}</p>
              <p className={`text-xs mt-1 ${
                message.type === 'user' ? 'text-primary-100' : 'text-gray-500'
              }`}>
                {message.timestamp.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
              </p>
            </div>
          </div>
        ))}
        
        {isTyping && (
          <div className="flex justify-start">
            <div className={`max-w-xs px-4 py-2 rounded-lg ${config.bgColor} ${config.color}`}>
              <div className="flex space-x-1">
                <div className="w-2 h-2 bg-current rounded-full animate-bounce"></div>
                <div className="w-2 h-2 bg-current rounded-full animate-bounce" style={{ animationDelay: '0.1s' }}></div>
                <div className="w-2 h-2 bg-current rounded-full animate-bounce" style={{ animationDelay: '0.2s' }}></div>
              </div>
            </div>
          </div>
        )}
        <div ref={messagesEndRef} />
      </div>

      {/* Input */}
      <div className="flex gap-2">
        <textarea
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
          onKeyPress={handleKeyPress}
          placeholder="Ask your buddy anything..."
          className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary-500 resize-none"
          rows={1}
          disabled={isTyping}
        />
        <button
          onClick={handleSendMessage}
          disabled={!inputValue.trim() || isTyping}
          className="btn-primary px-4 py-2 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          Send
        </button>
      </div>
    </div>
  );
}

