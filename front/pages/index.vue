<template>
  <p>{{ article }}</p>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted, useContext } from '@nuxtjs/composition-api'

export default defineComponent({
  name: 'TopPage',
  setup () {
    const article = ref(null)
    const { $axios } = useContext()

    onMounted(async () => {
      const { data } = await $axios.get('/api/sample')
      article.value = data
    })

    return {
      article
    }
  },
  head: {}
})
</script>
