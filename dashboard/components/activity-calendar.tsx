"use client";

import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { cn } from "@/lib/utils";
import { Calendar } from "lucide-react";
import { useMemo, useState } from "react";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "./ui/select";

interface ActivityData {
  date: string;
  count: number;
  level: 0 | 1 | 2 | 3 | 4;
  projects?: string[];
  editor?: string;
  details?: {
    commits?: number;
    pullRequests?: number;
    issues?: number;
    reviews?: number;
  };
}

function generateActivityData(period: string): ActivityData[] {
  const data: ActivityData[] = [];
  const today = new Date();
  let startDate: Date;

  if (period === "last12months") {
    startDate = new Date(today);
    startDate.setFullYear(today.getFullYear() - 1);
  } else {
    const year = Number.parseInt(period);
    startDate = new Date(year, 0, 1);
    const endDate =
      year === today.getFullYear() ? today : new Date(year, 11, 31);

    for (
      let d = new Date(startDate);
      d <= endDate;
      d.setDate(d.getDate() + 1)
    ) {
      data.push(generateDayActivity(new Date(d)));
    }
    return data;
  }

  for (let d = new Date(startDate); d <= today; d.setDate(d.getDate() + 1)) {
    data.push(generateDayActivity(new Date(d)));
  }

  return data;
}

function generateDayActivity(date: Date): ActivityData {
  const projects = ["Dashboard", "API", "Mobile App", "Website", "CLI Tool"];
  const editors = ["VS Code", "Cursor", "Vim", "WebStorm"];
  const dayOfWeek = date.getDay();
  const isWeekend = dayOfWeek === 0 || dayOfWeek === 6;

  let count = 0;
  let level: 0 | 1 | 2 | 3 | 4 = 0;

  if (Math.random() > (isWeekend ? 0.7 : 0.2)) {
    count = Math.floor(Math.random() * (isWeekend ? 3 : 8)) + 1;
    if (count >= 7) level = 4;
    else if (count >= 5) level = 3;
    else if (count >= 3) level = 2;
    else if (count >= 1) level = 1;
  }

  return {
    date: date.toISOString().split("T")[0],
    count,
    level,
    projects:
      count > 0 ? projects.slice(0, Math.floor(Math.random() * 3) + 1) : [],
    editor:
      count > 0
        ? editors[Math.floor(Math.random() * editors.length)]
        : undefined,
    details:
      count > 0
        ? {
            commits: Math.floor(Math.random() * count) + 1,
            pullRequests:
              Math.random() > 0.7 ? Math.floor(Math.random() * 2) + 1 : 0,
            issues: Math.random() > 0.8 ? Math.floor(Math.random() * 2) + 1 : 0,
            reviews:
              Math.random() > 0.6 ? Math.floor(Math.random() * 3) + 1 : 0,
          }
        : undefined,
  };
}

function getIntensityClass(level: 0 | 1 | 2 | 3 | 4): string {
  const intensity = {
    0: "bg-muted border border-border/50",
    1: "bg-primary/20 border border-primary/30",
    2: "bg-primary/40 border border-primary/50",
    3: "bg-primary/60 border border-primary/70",
    4: "bg-primary border border-primary",
  };

  return cn(
    "w-3 h-3 rounded-[2px] transition-all duration-200",
    "hover:ring-2 hover:ring-primary/50 hover:scale-110",
    intensity[level]
  );
}

function formatTooltipContent(activity: ActivityData): string {
  const date = new Date(activity.date).toLocaleDateString("en-US", {
    day: "numeric",
    month: "short",
  });

  if (activity.count === 0) {
    return `No activity on ${date}`;
  }

  return `${activity.count} activit${activity.count > 1 ? "ies" : "y"} on ${date}`;
}

export function ActivityCalendar() {
  const [selectedPeriod, setSelectedPeriod] = useState("last12months");
  const [selectedDay, setSelectedDay] = useState<ActivityData | null>(null);

  const availableYears = useMemo(() => {
    const currentYear = new Date().getFullYear();
    const startYear = 2020; // This would come from backend - when user started using the app
    const years = [];
    for (let year = currentYear; year >= startYear; year--) {
      years.push(year.toString());
    }
    return years;
  }, []);

  const activityData = useMemo(
    () => generateActivityData(selectedPeriod),
    [selectedPeriod]
  );

  // Group data by weeks for grid layout
  const weeks = useMemo(() => {
    const weeksArray: ActivityData[][] = [];
    let currentWeek: ActivityData[] = [];

    activityData.forEach((day, index) => {
      const dayOfWeek = new Date(day.date).getDay();

      // Start new week on Sunday
      if (dayOfWeek === 0 && currentWeek.length > 0) {
        weeksArray.push(currentWeek);
        currentWeek = [];
      }

      currentWeek.push(day);

      // Push the last week
      if (index === activityData.length - 1) {
        weeksArray.push(currentWeek);
      }
    });

    return weeksArray;
  }, [activityData]);

  const totalContributions = activityData.reduce(
    (sum, day) => sum + day.count,
    0
  );
  const currentStreak = useMemo(() => {
    let streak = 0;
    for (let i = activityData.length - 1; i >= 0; i--) {
      if (activityData[i].count > 0) {
        streak++;
      } else {
        break;
      }
    }
    return streak;
  }, [activityData]);

  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-4">
        <div className="flex flex-col gap-1">
          <CardTitle className="flex items-center gap-2 text-lg font-semibold">
            <Calendar className="text-primary h-5 w-5" />
            Activity Calendar
          </CardTitle>
          <p className="text-muted-foreground mt-1 text-sm">
            {totalContributions} contributions in the{" "}
            {selectedPeriod === "last12months" ? "last year" : selectedPeriod} •{" "}
            {currentStreak} day streak
          </p>
        </div>

        <Select value={selectedPeriod} onValueChange={setSelectedPeriod}>
          <SelectTrigger className="w-[180px]">
            <SelectValue placeholder="Select period" />
          </SelectTrigger>
          <SelectContent>
            <SelectItem value="last12months">Last 12 months</SelectItem>
            {availableYears.map((year) => (
              <SelectItem key={year} value={year}>
                {year}
              </SelectItem>
            ))}
          </SelectContent>
        </Select>
      </CardHeader>
      <CardContent>
        <div className="overflow-x-auto overflow-y-hidden">
          <div className="flex min-w-fit gap-1">
            {weeks.map((week, weekIndex) => (
              <div key={weekIndex} className="flex flex-col gap-1">
                {Array.from({ length: 7 }, (_, dayIndex) => {
                  const activity = week.find(
                    (day) => new Date(day.date).getDay() === dayIndex
                  );

                  if (!activity) {
                    return <div key={dayIndex} className="h-3 w-3" />;
                  }

                  return (
                    <Tooltip key={activity.date}>
                      <TooltipTrigger asChild>
                        <div
                          className={cn(
                            getIntensityClass(activity.level),
                            "cursor-pointer",
                            selectedDay?.date === activity.date &&
                              "ring-primary/70 scale-110 ring-2"
                          )}
                          role="button"
                          tabIndex={0}
                          onClick={() => setSelectedDay(activity)}
                          onKeyDown={(e) => {
                            if (e.key === "Enter" || e.key === " ") {
                              setSelectedDay(activity);
                            }
                          }}
                        />
                      </TooltipTrigger>
                      <TooltipContent side="top" className="max-w-xs">
                        <pre className="whitespace-pre-wrap text-xs">
                          {formatTooltipContent(activity)}
                        </pre>
                      </TooltipContent>
                    </Tooltip>
                  );
                })}
              </div>
            ))}
          </div>
        </div>
      </CardContent>
    </Card>
  );
}
