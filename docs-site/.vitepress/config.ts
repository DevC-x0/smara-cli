import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Smara CLI',
  description: 'Autonomous Multi-Agent Terminal with memory, skills, MCP, SSH, and knowledge graph tooling.',
  lang: 'id-ID',
  cleanUrls: true,
  lastUpdated: true,
  appearance: 'dark',
  head: [
    ['meta', { name: 'theme-color', content: '#0a0a0a' }],
    ['meta', { property: 'og:title', content: 'Smara CLI Documentation' }],
    ['meta', { property: 'og:description', content: 'Panduan resmi Smara CLI: install, quickstart, skills, memory, MCP, SSH, Graphify, dan automation.' }]
  ],
  themeConfig: {
    logo: '/logo.svg',
    siteTitle: 'Smara CLI',
    search: {
      provider: 'local'
    },
    nav: [
      { text: 'Getting Started', link: '/getting-started/installation' },
      { text: 'Concepts', link: '/core-concepts/ask-mode' },
      { text: 'Guides', link: '/guides/analyze-codebase' },
      { text: 'Reference', link: '/reference/cli-commands' },
      { text: 'Examples', link: '/examples/docs-generation' },
      { text: 'GitHub', link: 'https://github.com/gede-cahya/Smara-CLI' }
    ],
    sidebar: [
      {
        text: 'Getting Started',
        items: [
          { text: 'Introduction', link: '/' },
          { text: 'Installation', link: '/getting-started/installation' },
          { text: 'Quickstart', link: '/getting-started/quickstart' },
          { text: 'Configuration', link: '/getting-started/configuration' }
        ]
      },
      {
        text: 'Core Concepts',
        items: [
          { text: 'Ask Mode', link: '/core-concepts/ask-mode' },
          { text: 'Agent Mode', link: '/core-concepts/agent-mode' },
          { text: 'Memory', link: '/core-concepts/memory' },
          { text: 'Skills', link: '/core-concepts/skills' },
          { text: 'Tools', link: '/core-concepts/tools' },
          { text: 'MCP', link: '/core-concepts/mcp' },
          { text: 'Workflows', link: '/core-concepts/workflows' }
        ]
      },
      {
        text: 'Guides',
        items: [
          { text: 'Use VPS / SSH', link: '/guides/use-vps-ssh' },
          { text: 'Create a Skill', link: '/guides/create-skill' },
          { text: 'Create a Workflow', link: '/guides/create-workflow' },
          { text: 'Deploy a Project', link: '/guides/deploy-project' },
          { text: 'Analyze Codebase', link: '/guides/analyze-codebase' },
          { text: 'Knowledge Graph', link: '/guides/knowledge-graph' },
          { text: 'Generate Docs', link: '/guides/generate-docs' },
          { text: 'Provider Setup', link: '/guides/provider-setup' },
          { text: 'Release Process', link: '/guides/release-process' },
          { text: 'Web Interface', link: '/guides/web-interface' },
          { text: 'Platform Bots', link: '/guides/platform-bots' },
          { text: 'Evaluate Provider', link: '/guides/evaluate-provider' },
          { text: 'Browser Automation', link: '/guides/browser-automation' },
          { text: 'Multimodal Tools', link: '/guides/multimodal-tools' },
          { text: 'Scheduled Automation', link: '/guides/scheduled-automation' },
          { text: 'Workspaces', link: '/guides/workspaces' },
          { text: 'Cloud Memory', link: '/guides/cloud-memory' }
        ]
      },
      {
        text: 'Reference',
        items: [
          { text: 'CLI Commands', link: '/reference/cli-commands' },
          { text: 'Config Reference', link: '/reference/config' },
          { text: 'Tool List', link: '/reference/tool-list' },
          { text: 'Web API', link: '/reference/web-api' },
          { text: 'Graphify', link: '/reference/graphify' },
          { text: 'Skill Format', link: '/reference/skill-format' },
          { text: 'Analytics', link: '/reference/analytics' },
          { text: 'Sharing', link: '/reference/sharing' },
          { text: 'Policy', link: '/reference/policy' },
          { text: 'Run History', link: '/reference/run-history' },
          { text: 'Update', link: '/reference/update' },
          { text: 'Docs Gap Report', link: '/reference/docs-gap-report' },
          { text: 'Troubleshooting', link: '/reference/troubleshooting' }
        ]
      },
      {
        text: 'Examples',
        items: [
          { text: 'Docs Generation', link: '/examples/docs-generation' },
          { text: 'Docs Site Maintenance', link: '/examples/docs-site-maintenance' },
          { text: 'Provider Eval Report', link: '/examples/provider-eval-report' },
          { text: 'VPS Monitoring Skill', link: '/examples/vps-monitoring-skill' },
          { text: 'Release Checklist', link: '/examples/release-checklist' },
          { text: 'Cloud Memory Sync', link: '/examples/cloud-memory-sync' },
          { text: 'Browser E2E', link: '/examples/browser-e2e' },
          { text: 'React Deploy', link: '/examples/react-deploy' },
          { text: 'Golang Refactor', link: '/examples/golang-refactor' },
          { text: 'Server Monitoring', link: '/examples/server-monitoring' },
          { text: 'Common Workflows', link: '/examples/common-workflows' }
        ]
      }
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/gede-cahya/Smara-CLI' }
    ],
    footer: {
      message: 'Released under the MIT License.',
      copyright: '© 2026 Gede Cahya'
    },
    editLink: {
      pattern: 'https://github.com/gede-cahya/Smara-CLI/edit/main/docs-site/:path',
      text: 'Edit this page on GitHub'
    }
  }
})
