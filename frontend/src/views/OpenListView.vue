<script setup lang="ts">
import PageHeader from '@/components/common/PageHeader.vue'
</script>

<template>
  <section class="page">
    <PageHeader
      title="OpenList"
      description="Prepare this page for connection testing, driver listing, directory browsing, and raw-link validation."
    />

    <div class="placeholder-grid">
      <article class="placeholder-card">
        <h3>Connection Test</h3>
        <p>Reserve a form for endpoint, token, and connectivity checks.</p>
      </article>
      <article class="placeholder-card">
        <h3>Driver Catalog</h3>
        <p>Show the currently available OpenList drivers and adapter metadata.</p>
      </article>
      <article class="placeholder-card">
        <h3>Directory Browser</h3>
        <p>Plan for file listing, metadata inspection, and direct-link drilldown.</p>
      </article>
    </div>
  </section>
</template>
