<template>
  <section class="docs-graph" aria-label="Interactive documentation graph">
    <div class="docs-graph-toolbar">
      <div>
        <strong>Docs Graph</strong>
        <span>{{ visibleNodes.length }} pages · {{ visibleEdges.length }} links</span>
      </div>
      <div class="docs-graph-filters" role="tablist" aria-label="Filter docs graph">
        <button
          v-for="item in categories"
          :key="item"
          type="button"
          :class="{ active: filter === item }"
          @click="filter = item"
        >
          {{ item }}
        </button>
      </div>
    </div>

    <div class="docs-graph-stage">
      <svg viewBox="0 0 760 430" role="presentation" aria-hidden="true">
        <line
          v-for="edge in visibleEdges"
          :key="`${edge.from}-${edge.to}`"
          :x1="nodeById(edge.from)?.x"
          :y1="nodeById(edge.from)?.y"
          :x2="nodeById(edge.to)?.x"
          :y2="nodeById(edge.to)?.y"
        />
      </svg>

      <a
        v-for="node in visibleNodes"
        :key="node.id"
        class="docs-graph-node"
        :class="[`cat-${node.category}`, { selected: selected === node.id }]"
        :href="node.href"
        :style="{ left: `${node.x}px`, top: `${node.y}px` }"
        @mouseenter="selected = node.id"
        @focus="selected = node.id"
      >
        <span>{{ node.title }}</span>
        <small>{{ node.category }}</small>
      </a>
    </div>

    <p class="docs-graph-note">
      Graph ini ringan dan statis agar docs tetap cepat. Gunakan Graphify untuk graph source-code yang lebih besar.
    </p>
  </section>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'

type Category = 'start' | 'concept' | 'guide' | 'reference' | 'example'
type GraphNode = { id: string; title: string; href: string; category: Category; x: number; y: number }
type GraphEdge = { from: string; to: string }

const filter = ref<'all' | Category>('all')
const selected = ref<string>('smara')
const categories = ['all', 'start', 'concept', 'guide', 'reference', 'example'] as const

const nodes: GraphNode[] = [
  { id: 'smara', title: 'Smara CLI', href: '/', category: 'start', x: 330, y: 176 },
  { id: 'quickstart', title: 'Quickstart', href: '/getting-started/quickstart', category: 'start', x: 78, y: 62 },
  { id: 'memory', title: 'Memory', href: '/core-concepts/memory', category: 'concept', x: 290, y: 48 },
  { id: 'skills', title: 'Skills', href: '/core-concepts/skills', category: 'concept', x: 495, y: 74 },
  { id: 'mcp', title: 'MCP', href: '/core-concepts/mcp', category: 'concept', x: 590, y: 185 },
  { id: 'graphify', title: 'Graphify', href: '/reference/graphify', category: 'reference', x: 470, y: 306 },
  { id: 'webapi', title: 'Web API', href: '/reference/web-api', category: 'reference', x: 250, y: 308 },
  { id: 'provider', title: 'Provider Setup', href: '/guides/provider-setup', category: 'guide', x: 85, y: 190 },
  { id: 'docsgen', title: 'Generate Docs', href: '/guides/generate-docs', category: 'guide', x: 96, y: 306 },
  { id: 'examples', title: 'Examples', href: '/examples/common-workflows', category: 'example', x: 610, y: 312 }
]

const edges: GraphEdge[] = [
  { from: 'smara', to: 'quickstart' },
  { from: 'smara', to: 'memory' },
  { from: 'smara', to: 'skills' },
  { from: 'smara', to: 'mcp' },
  { from: 'smara', to: 'graphify' },
  { from: 'smara', to: 'webapi' },
  { from: 'smara', to: 'provider' },
  { from: 'smara', to: 'docsgen' },
  { from: 'smara', to: 'examples' },
  { from: 'docsgen', to: 'graphify' },
  { from: 'skills', to: 'examples' },
  { from: 'mcp', to: 'webapi' },
  { from: 'memory', to: 'graphify' }
]

const visibleNodes = computed(() => filter.value === 'all' ? nodes : nodes.filter(n => n.category === filter.value || n.id === 'smara'))
const visibleIds = computed(() => new Set(visibleNodes.value.map(n => n.id)))
const visibleEdges = computed(() => edges.filter(e => visibleIds.value.has(e.from) && visibleIds.value.has(e.to)))
const nodeById = (id: string) => nodes.find(n => n.id === id)
</script>
