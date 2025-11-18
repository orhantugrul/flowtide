"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Monitor } from "lucide-react";
import { Pie, PieChart } from "recharts";

const editorData = [
  {
    name: "VS Code",
    version: "1.85.2",
    timeSpent: 385,
    percentage: 68,
    sessions: 12,
    avgSessionTime: 32,
    fill: "#007acc", //vs code color
    trend: "+5%",
    status: "active",
  },
  {
    name: "Cursor",
    version: "0.29.1",
    timeSpent: 120,
    percentage: 21,
    sessions: 4,
    avgSessionTime: 30,
    fill: "#000000", //cursor color
    trend: "+15%",
    status: "recent",
  },
  {
    name: "IntelliJ IDEA",
    version: "2.1.0",
    timeSpent: 80,
    percentage: 14,
    sessions: 3,
    avgSessionTime: 26,
    fill: "#ff7a00", //intellij idea color
    trend: "+10%",
    status: "active",
  },
];

const totalTime = editorData.reduce((sum, editor) => sum + editor.timeSpent, 0);

const chartConfig = {
  "VS Code": { label: "VS Code" },
  Cursor: { label: "Cursor" },
  "IntelliJ IDEA": { label: "IntelliJ IDEA" },
};

export function EditorUsage() {
  return (
    <Card>
      <CardHeader className="pb-4">
        <CardTitle className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Monitor className="text-primary h-5 w-5" />
            Editor Usage
          </div>
          <div className="text-muted-foreground text-sm font-normal">
            {Math.floor(totalTime / 60)}h {totalTime % 60}m total
          </div>
        </CardTitle>
      </CardHeader>
      <CardContent className="flex-1 pb-0">
        <ChartContainer
          config={chartConfig}
          className="mx-auto aspect-square max-h-[300px]"
        >
          <PieChart>
            <ChartTooltip
              content={<ChartTooltipContent nameKey="visitors" hideLabel />}
            />
            <Pie
              data={editorData}
              dataKey="timeSpent"
              stroke="var(--background)"
              fillOpacity={0.8}
            ></Pie>
            <ChartLegend
              content={<ChartLegendContent nameKey="name" />}
              className="-translate-y-2 flex-wrap gap-2 *:basis-1/4 *:justify-center"
            />
          </PieChart>
        </ChartContainer>
      </CardContent>
    </Card>
  );
}
