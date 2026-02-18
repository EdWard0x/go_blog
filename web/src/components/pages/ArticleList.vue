<template>
  <div class="article-list">
    <div class="list-header">
      <h2 class="section-title">文章列表</h2>
      <div class="search">
        <el-input v-model="articleSearchRequest.query" placeholder="搜索文章..." prefix-icon="Search" maxlength="50"
                  @change="changeArticleSearchItem" clearable/>
        <el-button type="primary" @click="changeArticleSearchItem">搜索</el-button>
      </div>
    </div>

    <div class="filter-section">
      <div class="filter-group">
        <span class="filter-label">类别</span>
        <el-radio-group v-model="articleSearchRequest.category" @change="changeArticleSearchItem" size="default">
          <el-radio-button label="全部" value=""/>
          <template v-for="item in categoryArr">
            <el-radio-button :label="item" :value="item"/>
          </template>
        </el-radio-group>
      </div>

      <div class="filter-group">
        <span class="filter-label">标签</span>
        <el-radio-group v-model="articleSearchRequest.tag" @change="changeArticleSearchItem" size="default">
          <el-radio-button label="全部" value=""/>
          <template v-for="item in tagArr">
            <el-radio-button :label="item" :value="item"/>
          </template>
        </el-radio-group>
      </div>

      <div class="filter-group sort-group">
        <span class="filter-label">排序</span>
        <el-button-group>
          <el-button @click="handleSortClick();changeArticleSearchItem()" :type="articleSearchRequest.order === 'desc' ? 'primary' : 'default'">
            <el-icon><component is="SortDown"/></el-icon>
          </el-button>
          <el-button @click="handleSortClick();changeArticleSearchItem()" :type="articleSearchRequest.order === 'asc' ? 'primary' : 'default'">
            <el-icon><component is="SortUp"/></el-icon>
          </el-button>
        </el-button-group>
        <el-radio-group v-model="articleSearchRequest.sort" @change="changeArticleSearchItem" size="default">
          <el-radio-button label="默认" value=""/>
          <el-radio-button label="时间" value="time"/>
          <el-radio-button label="评论" value="comment"/>
          <el-radio-button label="浏览" value="view"/>
          <el-radio-button label="点赞" value="like"/>
        </el-radio-group>
      </div>
    </div>

    <div class="articles-container">
      <div 
        v-for="article in articleTableData" 
        :key="article._id" 
        class="article-card"
        @click="handleArticleJumps(article._id)"
      >
        <div class="article-cover">
          <el-image fit="cover" :src="article._source.cover" alt="">
            <template #error>
              <div class="image-placeholder">
                <el-icon><Picture /></el-icon>
              </div>
            </template>
          </el-image>
        </div>
        <div class="article-content">
          <h3 class="article-title">{{ article._source.title }}</h3>
          <p class="article-abstract">{{ article._source.abstract }}</p>
          <div class="article-footer">
            <div class="article-tags">
              <el-tag v-for="item in article._source.tags?.slice(0, 3)" :key="item" size="small" effect="plain">
                {{ item }}
              </el-tag>
            </div>
            <div class="article-meta">
              <span class="meta-item">
                <el-icon><Calendar /></el-icon>
                {{ article._source.created_at }}
              </span>
              <span class="meta-item">
                <el-icon><View /></el-icon>
                {{ article._source.views }}
              </span>
              <span class="meta-item">
                <el-icon><ChatDotRound /></el-icon>
                {{ article._source.comments }}
              </span>
              <span class="meta-item">
                <el-icon><Star /></el-icon>
                {{ article._source.likes }}
              </span>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="pagination-wrapper">
      <el-pagination
          :current-page="page"
          :page-size="page_size"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
          background
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import type {Hit} from "@/api/common";
import {type Article, articleCategory, articleSearch, type ArticleSearchRequest, articleTags} from "@/api/article";
import {computed, reactive, ref} from "vue";
import {Calendar, Picture, View, ChatDotRound, Star} from '@element-plus/icons-vue';

const articleSearchRequest = reactive<ArticleSearchRequest>({
  query: "",
  category: "",
  tag: "",
  sort: "",
  order: "desc",
  page: 1,
  page_size: 10,
})

const categoryArr = ref<string[]>([])
const tagArr = ref<string[]>([])

const handleSortClick = () => {
  articleSearchRequest.order = articleSearchRequest.order === "desc" ? "asc" : "desc"
}

const getArticleCategory = async () => {
  const res = await articleCategory()
  if (res.code === 0) {
    res.data.forEach((item) => {
      categoryArr.value.push(item.category)
    })
  }
}

getArticleCategory()

const getArticleTags = async () => {
  const res = await articleTags()
  if (res.code === 0) {
    res.data.forEach((item) => {
      tagArr.value.push(item.tag)
    })
  }
}

getArticleTags()

const page = ref(1)
const page_size = ref(10)
const total = ref(0)
const articleTableData = ref<Hit<Article>[]>()

const getArticleSearchTableData = async () => {
  articleSearchRequest.page = page.value;
  articleSearchRequest.page_size = page_size.value;

  const table = await articleSearch(articleSearchRequest)

  if (table.code === 0) {
    articleTableData.value = table.data.list;
    total.value = table.data.total;
  }
}

getArticleSearchTableData()

const changeArticleSearchItem = () => {
  getArticleSearchTableData()
}

const handleArticleJumps = (id: string) => {
  window.open("/article/" + id)
}

const handleSizeChange = (val: number) => {
  page_size.value = val
  getArticleSearchTableData()
}

const handleCurrentChange = (val: number) => {
  page.value = val
  getArticleSearchTableData()
}
</script>

<style scoped lang="scss">
.article-list {
  background: var(--bg-card);
  border-radius: var(--radius-lg);
  padding: 30px;
  box-shadow: var(--shadow-sm);
  
  .list-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 24px;
    padding-bottom: 20px;
    border-bottom: 1px solid var(--border-color-light);
    
    .section-title {
      font-size: 1.75rem;
      font-weight: 600;
      color: var(--text-primary);
      margin: 0;
      position: relative;
      
      &::after {
        content: '';
        position: absolute;
        bottom: -20px;
        left: 0;
        width: 50px;
        height: 3px;
        background: var(--color-primary);
        border-radius: 2px;
      }
    }
    
    .search {
      display: flex;
      gap: 12px;
      
      .el-input {
        width: 280px;
      }
    }
  }
  
  .filter-section {
    margin-bottom: 24px;
    padding: 20px;
    background: var(--bg-secondary);
    border-radius: var(--radius-md);
    
    .filter-group {
      display: flex;
      align-items: center;
      flex-wrap: wrap;
      gap: 12px;
      margin-bottom: 16px;
      
      &:last-child {
        margin-bottom: 0;
      }
      
      .filter-label {
        font-weight: 500;
        color: var(--text-secondary);
        min-width: 50px;
      }
      
      .el-radio-group {
        flex-wrap: wrap;
        gap: 8px;
        
        :deep(.el-radio-button__inner) {
          border-radius: var(--radius-sm);
          border: 1px solid var(--border-color);
          background: var(--bg-card);
          color: var(--text-secondary);
          transition: all 0.2s ease;
          
          &:hover {
            border-color: var(--color-primary);
            color: var(--color-primary);
          }
        }
        
        :deep(.el-radio-button.is-active .el-radio-button__inner) {
          background: var(--color-primary);
          border-color: var(--color-primary);
          color: #fff;
          box-shadow: none;
        }
      }
    }
    
    .sort-group {
      .el-button-group {
        margin-right: 12px;
        
        .el-button {
          padding: 8px 12px;
        }
      }
    }
  }
  
  .articles-container {
    display: flex;
    flex-direction: column;
    gap: 20px;
  }
  
  .article-card {
    display: flex;
    gap: 24px;
    padding: 20px;
    background: var(--bg-card);
    border-radius: var(--radius-md);
    border: 1px solid var(--border-color-light);
    cursor: pointer;
    transition: all 0.3s ease;
    
    &:hover {
      border-color: var(--color-primary);
      box-shadow: var(--shadow-md);
      transform: translateY(-2px);
      
      .article-cover {
        .el-image {
          transform: scale(1.05);
        }
      }
      
      .article-title {
        color: var(--color-primary);
      }
    }
    
    .article-cover {
      flex-shrink: 0;
      width: 200px;
      height: 130px;
      border-radius: var(--radius-sm);
      overflow: hidden;
      
      .el-image {
        width: 100%;
        height: 100%;
        transition: transform 0.3s ease;
      }
      
      .image-placeholder {
        width: 100%;
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: center;
        background: var(--bg-secondary);
        color: var(--text-muted);
        font-size: 2rem;
      }
    }
    
    .article-content {
      flex: 1;
      display: flex;
      flex-direction: column;
      min-width: 0;
      
      .article-title {
        font-size: 1.25rem;
        font-weight: 600;
        color: var(--text-primary);
        margin: 0 0 8px 0;
        transition: color 0.2s ease;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
      
      .article-abstract {
        flex: 1;
        color: var(--text-secondary);
        font-size: 0.95rem;
        line-height: 1.6;
        margin: 0 0 12px 0;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        overflow: hidden;
      }
      
      .article-footer {
        display: flex;
        justify-content: space-between;
        align-items: center;
        flex-wrap: wrap;
        gap: 12px;
        
        .article-tags {
          display: flex;
          gap: 8px;
          flex-wrap: wrap;
          
          .el-tag {
            background: rgba(90, 159, 143, 0.1);
            border-color: rgba(90, 159, 143, 0.2);
            color: var(--color-primary);
          }
        }
        
        .article-meta {
          display: flex;
          gap: 16px;
          color: var(--text-muted);
          font-size: 0.85rem;
          
          .meta-item {
            display: flex;
            align-items: center;
            gap: 4px;
            
            .el-icon {
              font-size: 1rem;
            }
          }
        }
      }
    }
  }
  
  .pagination-wrapper {
    margin-top: 30px;
    display: flex;
    justify-content: center;
    
    :deep(.el-pagination) {
      .el-pager li {
        background: var(--bg-card);
        border: 1px solid var(--border-color);
        border-radius: var(--radius-sm);
        margin: 0 4px;
        
        &.is-active {
          background: var(--color-primary);
          border-color: var(--color-primary);
        }
      }
      
      .btn-prev,
      .btn-next {
        background: var(--bg-card);
        border: 1px solid var(--border-color);
        border-radius: var(--radius-sm);
      }
    }
  }
}

@media (max-width: 768px) {
  .article-list {
    padding: 20px;
    
    .list-header {
      flex-direction: column;
      align-items: flex-start;
      gap: 16px;
      
      .search {
        width: 100%;
        
        .el-input {
          flex: 1;
          width: auto;
        }
      }
    }
    
    .article-card {
      flex-direction: column;
      
      .article-cover {
        width: 100%;
        height: 180px;
      }
    }
  }
}
</style>
