import { ActiveProjects } from "@/components/active-projects";
import { ActivityCalendar } from "@/components/activity-calendar";
import { ActivityInsights } from "@/components/activity-insights";
import { AvgDaily } from "@/components/average-daily";
import { DailyActivity } from "@/components/charts/daily-activity";
import { EditorUsage } from "@/components/charts/editor-usage";
import { LanguageUsage } from "@/components/charts/language-usage";
import { TotalHours } from "@/components/total-hours";

export default function Dashboard() {
  return (
    <div className="flex flex-col">
      <div className="mb-6 grid grid-cols-1 gap-4 md:grid-cols-3">
        <TotalHours />
        <ActiveProjects />
        <AvgDaily />
      </div>
      <div className="mb-6">
        <DailyActivity />
      </div>
      <div className="mb-6 grid grid-cols-1 gap-4 md:grid-cols-2 md:gap-6">
        <LanguageUsage />
        <EditorUsage />
      </div>
      <div className="grid grid-cols-1 gap-6 lg:grid-cols-4">
        <div className="lg:col-span-3">
          <ActivityCalendar />
        </div>
        <div className="lg:col-span-1">
          <ActivityInsights />
        </div>
      </div>
    </div>
  );
}
