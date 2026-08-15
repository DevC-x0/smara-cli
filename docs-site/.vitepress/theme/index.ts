import DefaultTheme from 'vitepress/theme'
import DocsGraph from './components/DocsGraph.vue'
import './custom.css'

export default {
  extends: DefaultTheme,
  enhanceApp({ app }) {
    app.component('DocsGraph', DocsGraph)
  }
}
