"use client";

import { useMemo } from "react";

interface ActivityData {
  date: string;
  count: number;
  level: 0 | 1 | 2 | 3 | 4;
  projects?: string[];
  editor?: string;
}

// Generate mock activity data for the past year
function generateActivityData(): ActivityData[] {
  const data: ActivityData[] = [];
  const today = new Date();
  const oneYearAgo = new Date(today);
  oneYearAgo.setFullYear(today.getFullYear() - 1);

  const projects = ["Dashboard", "API", "Mobile App", "Website", "CLI Tool"];
  const editors = ["VS Code", "Cursor", "Vim", "WebStorm"];

  for (let d = new Date(oneYearAgo); d <= today; d.setDate(d.getDate() + 1)) {
    const dayOfWeek = d.getDay();
    const isWeekend = dayOfWeek === 0 || dayOfWeek === 6;

    // Generate realistic activity patterns
    let count = 0;
    let level: 0 | 1 | 2 | 3 | 4 = 0;

    if (Math.random() > (isWeekend ? 0.7 : 0.2)) {
      count = Math.floor(Math.random() * (isWeekend ? 3 : 8)) + 1;
      if (count >= 7) level = 4;
      else if (count >= 5) level = 3;
      else if (count >= 3) level = 2;
      else if (count >= 1) level = 1;
    }

    data.push({
      date: d.toISOString().split("T")[0],
      count,
      level,
      projects:
        count > 0 ? projects.slice(0, Math.floor(Math.random() * 3) + 1) : [],
      editor:
        count > 0
          ? editors[Math.floor(Math.random() * editors.length)]
          : undefined,
    });
  }

  return data;
}

export function ActivityInsights() {
  const activityData = useMemo(() => generateActivityData(), []);

  const insights = useMemo(() => {
    const totalContributions = activityData.reduce(
      (sum, day) => sum + day.count,
      0
    );
    const activeDays = activityData.filter((day) => day.count > 0).length;
    const totalDays = activityData.length;

    // Calculate current streak
    let currentStreak = 0;
    for (let i = activityData.length - 1; i >= 0; i--) {
      if (activityData[i].count > 0) {
        currentStreak++;
      } else {
        break;
      }
    }

    // Calculate longest streak
    let longestStreak = 0;
    let tempStreak = 0;
    activityData.forEach((day) => {
      if (day.count > 0) {
        tempStreak++;
        longestStreak = Math.max(longestStreak, tempStreak);
      } else {
        tempStreak = 0;
      }
    });

    // Find most productive day
    const mostProductiveDay = activityData.reduce(
      (max, day) => (day.count > max.count ? day : max),
      activityData[0] || { count: 0, date: "" }
    );

    // Calculate weekly average
    const weeklyAverage =
      Math.round((totalContributions / totalDays) * 7 * 10) / 10;

    // Activity consistency (percentage of days with activity)
    const consistency = Math.round((activeDays / totalDays) * 100);

    return {
      totalContributions,
      currentStreak,
      longestStreak,
      mostProductiveDay,
      weeklyAverage,
      consistency,
      activeDays,
    };
  }, [activityData]);

  return (
    <div className="space-y-4">
      {}
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground text-sm">Current Streak</span>
        <span className="text-sm font-medium">{insights.currentStreak}</span>
      </div>
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground text-sm">Best Streak</span>
        <span className="text-sm font-medium">{insights.longestStreak}</span>
      </div>
      {insights.mostProductiveDay.count > 0 && (
        <div className="flex items-center justify-between">
          <span className="text-muted-foreground text-sm">Best Day</span>
          <span className="text-sm font-medium">
            {insights.mostProductiveDay.count}
          </span>
        </div>
      )}
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground text-sm">Weekly Average</span>
        <span className="text-sm font-medium">{insights.weeklyAverage}</span>
      </div>
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground text-sm">Consistency</span>
        <span className="text-sm font-medium">{insights.consistency}%</span>
      </div>
      <div className="flex items-center justify-between">
        <span className="text-muted-foreground text-sm">Active Days</span>
        <span className="text-sm font-medium">{insights.activeDays}</span>
      </div>
    </div>
  );
}
