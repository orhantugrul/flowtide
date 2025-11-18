"use client";

import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import {
  ChartContainer,
  ChartLegend,
  ChartLegendContent,
  ChartTooltip,
  ChartTooltipContent,
} from "@/components/ui/chart";
import { Code2 } from "lucide-react";
import { Pie, PieChart } from "recharts";

const languageData = [
  {
    name: "TypeScript",
    hours: 452,
    percentage: 38,
    fill: "#007acc", //type script color
    trend: "up",
    sessions: 28,
  },
  {
    name: "JavaScript",
    hours: 321,
    percentage: 27,
    fill: "#ffd700", //javascript yellow
    trend: "down",
    sessions: 19,
  },
  {
    name: "Python",
    hours: 248,
    percentage: 21,
    fill: "#3572a5", //python blue
    trend: "up",
    sessions: 15,
  },
  {
    name: "Go",
    hours: 105,
    percentage: 9,
    fill: "#00add8", //go blue
    trend: "up",
    sessions: 8,
  },
  {
    name: "Rust",
    hours: 62,
    percentage: 5,
    fill: "#b7410e", //rust orange
    trend: "up",
    sessions: 4,
  },
];

const totalTime = languageData.reduce(
  (sum, language) => sum + language.hours,
  0
);

const chartConfig = {
  TypeScript: { label: "TypeScript" },
  JavaScript: { label: "JavaScript" },
  Python: { label: "Python" },
  Go: { label: "Go" },
  Rust: { label: "Rust" },
};

export function LanguageUsage() {
  return (
    <Card>
      <CardHeader className="pb-4">
        <CardTitle className="flex items-center justify-between">
          <div className="flex items-center gap-2">
            <Code2 className="text-primary h-5 w-5" />
            Language Usage
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
            <ChartTooltip content={<ChartTooltipContent nameKey="hours" />} />
            <Pie
              data={languageData}
              dataKey="hours"
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
