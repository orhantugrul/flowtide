"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { TrendingUp } from "lucide-react";
import { useState } from "react";
import { Area, AreaChart, CartesianGrid, XAxis } from "recharts";
import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from "../ui/chart";

const data = [
  { date: "2025-04-01", hours: 4.2, sessions: 8 },
  { date: "2025-04-02", hours: 6.8, sessions: 12 },
  { date: "2025-04-03", hours: 3.1, sessions: 6 },
  { date: "2025-04-04", hours: 8.5, sessions: 15 },
  { date: "2025-04-05", hours: 5.9, sessions: 11 },
  { date: "2025-04-06", hours: 2.3, sessions: 4 },
  { date: "2025-04-07", hours: 7.2, sessions: 13 },
  { date: "2025-04-08", hours: 9.1, sessions: 18 },
  { date: "2025-04-09", hours: 4.7, sessions: 9 },
  { date: "2025-04-10", hours: 6.3, sessions: 12 },
  { date: "2025-04-11", hours: 3.8, sessions: 7 },
  { date: "2025-04-12", hours: 8.9, sessions: 16 },
  { date: "2025-04-13", hours: 5.4, sessions: 10 },
  { date: "2025-04-14", hours: 7.6, sessions: 14 },
  { date: "2025-04-15", hours: 4.1, sessions: 8 },
  { date: "2025-04-16", hours: 6.7, sessions: 13 },
  { date: "2025-04-17", hours: 8.3, sessions: 15 },
  { date: "2025-04-18", hours: 5.2, sessions: 9 },
  { date: "2025-04-19", hours: 7.8, sessions: 14 },
  { date: "2025-04-20", hours: 9.4, sessions: 17 },
  { date: "2025-04-21", hours: 6.1, sessions: 11 },
  { date: "2025-04-22", hours: 4.9, sessions: 8 },
  { date: "2025-04-23", hours: 3.2, sessions: 6 },
  { date: "2025-04-24", hours: 1.8, sessions: 3 },
  { date: "2025-04-25", hours: 0.5, sessions: 1 },
  { date: "2025-04-26", hours: 2.1, sessions: 4 },
  { date: "2025-04-27", hours: 5.7, sessions: 10 },
  { date: "2025-04-28", hours: 8.2, sessions: 15 },
  { date: "2025-04-29", hours: 7.3, sessions: 13 },
  { date: "2025-04-30", hours: 6.8, sessions: 12 },
];

const chartConfig = {
  visitors: {
    label: "Visitors",
  },
  hours: {
    label: "Hours",
    color: "var(--chart-1)",
  },
} satisfies ChartConfig;

export function DailyActivity() {
  const [timeRange, setTimeRange] = useState<"7days" | "30days">("30days");

  const filteredData = data.filter((item) => {
    const date = new Date(item.date);
    const referenceDate = new Date("2025-04-30");
    const ranges = { "7days": 7, "30days": 30 };

    const startDate = new Date(referenceDate);
    startDate.setDate(startDate.getDate() - ranges[timeRange]);
    return date >= startDate;
  });

  return (
    <Card>
      <CardHeader className="pb-4">
        <div className="flex items-center justify-between">
          <CardTitle className="flex items-center gap-2 text-lg font-semibold">
            <TrendingUp className="text-primary h-5 w-5" />
            Daily Activity
          </CardTitle>
          <Select
            value={timeRange}
            onValueChange={(value: "7days" | "30days") => setTimeRange(value)}
          >
            <SelectTrigger className="h-8 w-32 text-xs">
              <SelectValue />
            </SelectTrigger>
            <SelectContent>
              <SelectItem value="7days">Last 7 Days</SelectItem>
              <SelectItem value="30days">Last 30 days</SelectItem>
            </SelectContent>
          </Select>
        </div>
      </CardHeader>
      <CardContent>
        <ChartContainer
          config={chartConfig}
          className="aspect-auto h-[250px] w-full"
        >
          <AreaChart data={filteredData}>
            <defs>
              <linearGradient id="fillHours" x1="0" y1="0" x2="0" y2="1">
                <stop
                  offset="5%"
                  stopColor="var(--color-hours)"
                  stopOpacity={0.8}
                />
                <stop
                  offset="95%"
                  stopColor="var(--color-hours)"
                  stopOpacity={0.1}
                />
              </linearGradient>
            </defs>
            <CartesianGrid vertical={false} />
            <XAxis
              dataKey="date"
              tickLine={false}
              axisLine={false}
              tickMargin={8}
              minTickGap={32}
              tickFormatter={(value) => {
                const date = new Date(value);
                return date.toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                });
              }}
            />
            <ChartTooltip
              cursor={false}
              content={
                <ChartTooltipContent
                  labelFormatter={(value) => {
                    return new Date(value).toLocaleDateString("en-US", {
                      month: "short",
                      day: "numeric",
                    });
                  }}
                  indicator="dot"
                />
              }
            />
            <Area
              dataKey="hours"
              type="natural"
              fill="url(#fillHours)"
              stroke="var(--color-hours)"
              stackId="a"
            />
          </AreaChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
