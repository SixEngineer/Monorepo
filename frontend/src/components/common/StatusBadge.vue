<script setup lang="ts">
import { computed } from 'vue'
import type { HealthState } from '@/types/common'

const props = defineProps<{
  state: HealthState
}>()

const label = computed(() => {
  switch (props.state) {
    case 'active':
      return 'Active'
    case 'disabled':
      return 'Disabled'
    case 'expired':
      return 'Expired'
    case 'error':
      return 'Error'
    default:
      return props.state
  }
})
</script>

<template>
  <span class="status-badge" :class="`status-badge--${state}`">{{ label }}</span>
</template>