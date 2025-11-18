import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Clock } from "lucide-react";

export function TotalHours() {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">Total Hours</CardTitle>
        <Clock className="text-muted-foreground h-4 w-4" />
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">127.5h</div>
        <p className="text-muted-foreground text-xs">+12% from last month</p>
      </CardContent>
    </Card>
  );
}
