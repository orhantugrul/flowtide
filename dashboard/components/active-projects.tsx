import { Code } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "./ui/card";

export function ActiveProjects() {
  return (
    <Card>
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium">Active Projects</CardTitle>
        <Code className="text-muted-foreground h-4 w-4" />
      </CardHeader>
      <CardContent>
        <div className="text-2xl font-bold">8</div>
        <p className="text-muted-foreground text-xs">+2 new this week</p>
      </CardContent>
    </Card>
  );
}
