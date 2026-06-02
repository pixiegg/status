<template>
  <Card class="overflow-hidden">
    <CardContent class="p-6 flex items-center gap-4">
      <div
        class="relative flex h-6 w-6 items-center justify-center flex-shrink-0"
      >
        <span
          :class="[
            'absolute inline-flex h-full w-full rounded-full opacity-75 animate-ping',
            dotColorClass,
          ]"
        ></span>
        <span
          :class="['relative inline-flex rounded-full h-4 w-4', dotColorClass]"
        ></span>
      </div>
      <h2 class="text-2xl sm:text-3xl font-bold tracking-tight text-foreground">
        {{ statusText }}
      </h2>
    </CardContent>
  </Card>
</template>

<script setup>
import { computed } from "vue";
import { Card, CardContent } from "@/components/ui/card";

const props = defineProps({
  endpoints: {
    type: Array,
    required: true,
    default: () => [],
  },
});

const getEndpointStatus = (endpoint) => {
  if (
    endpoint.status &&
    ["healthy", "unhealthy", "degraded"].includes(endpoint.status)
  ) {
    return endpoint.status;
  }
  if (!endpoint.results || endpoint.results.length === 0) return "unknown";
  const latestResult = endpoint.results[endpoint.results.length - 1];
  return latestResult.success ? "healthy" : "unhealthy";
};

const overallState = computed(() => {
  if (!props.endpoints || props.endpoints.length === 0) {
    return { status: "unknown", text: "Aguardando status..." };
  }

  const statuses = props.endpoints
    .map(getEndpointStatus)
    .filter((s) => s !== "unknown");

  if (statuses.length === 0) {
    return { status: "unknown", text: "Aguardando status..." };
  }

  const allHealthy = statuses.every((s) => s === "healthy");
  if (allHealthy) {
    return { status: "healthy", text: "Todos sistemas operando" };
  }

  const allUnhealthy = statuses.every((s) => s === "unhealthy");
  if (allUnhealthy) {
    return { status: "unhealthy", text: "Indisponibilidade total" };
  }

  const allDegraded = statuses.every((s) => s === "degraded");
  if (allDegraded) {
    return { status: "degraded", text: "Lentidão total" };
  }

  const anyUnhealthy = statuses.some((s) => s === "unhealthy");
  if (anyUnhealthy) {
    return { status: "unhealthy_partial", text: "Indisponibilidade parcial" };
  }

  const anyDegraded = statuses.some((s) => s === "degraded");
  if (anyDegraded) {
    return { status: "degraded_partial", text: "Lentidão parcial" };
  }

  return { status: "unknown", text: "Status desconhecido" };
});

const statusText = computed(() => overallState.value.text);

const dotColorClass = computed(() => {
  switch (overallState.value.status) {
    case "healthy":
      return "bg-green-400";
    case "unhealthy":
    case "unhealthy_partial":
      return "bg-red-400";
    case "degraded":
    case "degraded_partial":
      return "bg-yellow-400";
    default:
      return "bg-gray-400";
  }
});
</script>
