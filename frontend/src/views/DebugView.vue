<script setup lang="ts">
import PageHeader from '@/components/common/PageHeader.vue'
</script>

<template>
  <section class="page">
    <PageHeader
      title="Debug"
      description="Reserve this space for redirect chains, final URLs, headers, cache-hit markers, and provider trace output."
    />

    <section class="panel debug-grid">
      <article class="placeholder-card">
        <h3>Resolver Trace</h3>
        <p>redirect_chain, final_url, fallback_reason</p>
      </article>
      <article class="placeholder-card">
        <h3>Provider Debug</h3>
        <p>Capability checks, auth status, and recent diagnostics.</p>
      </article>
      <article class="placeholder-card">
        <h3>Task Debug</h3>
        <p>Headers, request context, retry signals, and transport notes.</p>
      </article>
    </section>
  </section>
</template>
