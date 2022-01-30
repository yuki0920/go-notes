<template>
  <div class="container">
    <div class="row">
      <!-- TODO: 入力フォームと共通化したい -->
      <form>
        <!-- TODO: TODO: エラーを表示したい -->
        <div>
          <label for="form-title">タイトル</label>
          <input id="form-title" v-model=" article.title" type="text" name="title">
        </div>
        <div>
          <multiselect
            v-model="selectedTags"
            tag-placeholder="Add this as new tag"
            placeholder="Search or add a tag"
            label="title"
            track-by="id"
            :options="categoryOptions"
            :multiple="true"
            :taggable="true"
            @tag="addTag"
          />
        </div>
        <div>
          <label for="form-body">本文</label>
          <textarea
            id="form-body"
            v-model="article.body"
            rows="20"
            name="body"
          />
        </div>

        <div>
          <a href="javascript:void(0)" class="btn btn-secondary">
            <nuxt-link to="/">
              キャンセル
            </nuxt-link>
          </a>
          <a href="javascript:void(0)" class="btn btn-dark" @click="submit">
            保存
          </a>
        </div>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, useRouter, useContext, reactive, ref, onMounted } from '@nuxtjs/composition-api'
import { Category } from '~/types/category'

export default defineComponent({
  name: 'ArticleNew',
  setup () {
    const { $axios } = useContext()
    const article = reactive({ title: null, body: null })
    const router = useRouter()

    type categoryOption = {
      id: number,
      title: string,
    }
    const categoryOptions = ref<categoryOption[]>([])
    const selectedTags = ref<categoryOption[]>([])

    onMounted(async () => {
      const { data: categoryData } = await $axios.get('/api/categories')
      if (categoryData) {
        categoryData.forEach((category: Category) => {
          categoryOptions.value.push({ id: category.id, title: category.title })
        })
      }
    })

    const addTag = async (newTag: string) => {
      const { data: id } = await $axios.post('/api/categories', { title: newTag })
      const tag = {
        id: parseInt(id, 10),
        title: newTag
      }
      categoryOptions.value.push(tag)
      selectedTags.value.push(tag)
    }

    const submit = async (event: any) => {
      event.preventDefault()
      const params = {
        title: article.title,
        body: article.body,
        category_ids: selectedTags.value.map(tag => tag.id)
      }
      const { data } = await $axios.post('/api/articles', params)
      const id = data.Article.id
      router.push(`/articles/${id}`)
    }

    return {
      article,
      submit,
      categoryOptions,
      selectedTags,
      addTag
    }
  }
})
</script>

<style lang='scss' scoped>
</style>
