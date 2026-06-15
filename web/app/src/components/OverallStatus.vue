<template>
  <Card class="overflow-hidden">
    <CardContent class="p-4 sm:p-5 flex items-center gap-3">
      <div
        class="relative flex h-5 w-5 items-center justify-center flex-shrink-0"
      >
        <span
          :class="[
            'absolute inline-flex h-full w-full rounded-full opacity-75 animate-ping',
            dotColorClass,
          ]"
        ></span>
        <span
          :class="['relative inline-flex rounded-full h-3 w-3', dotColorClass]"
        ></span>
      </div>
      <h2
        class="ml-2 text-lg sm:text-xl font-medium tracking-tight text-foreground"
      >
        {{ statusText }}
      </h2>
    </CardContent>
  </Card>
</template>

<script setup>
import { computed } from "vue";
import { Card, CardContent } from "@/components/ui/card";

const props = defineProps({
  status: {
    type: String,
    required: true,
    default: "unknown",
  },
});

const statusMap = {
  healthy: "Todos sistemas operando",
  unhealthy: "Indisponibilidade total",
  degraded: "Lentidão total",
  unhealthy_partial: "Indisponibilidade parcial",
  degraded_partial: "Lentidão parcial",
};

const statusText = computed(() => {
  return statusMap[props.status] || "Status desconhecido";
});

const dotColorClass = computed(() => {
  switch (props.status) {
    case "healthy":
      return "bg-green-400";
    case "unhealthy":
    case "degraded":
      return "bg-red-400";
    case "unhealthy_partial":
    case "degraded_partial":
      return "bg-yellow-400";
    default:
      return "bg-gray-400";
  }
});
</script>
