<template>
  <el-card class="tag-cloud">
    <div class="section-header">
      <el-icon class="header-icon"><PriceTag /></el-icon>
      <span class="title">标签云</span>
    </div>
    <div class="tag-container">
      <span 
        v-for="item in tagCloudArray" 
        :key="item.tag" 
        :class="['tag-item', `tag-${item.type}`]"
        @click="handleSearchJumps(item.tag)"
      >
        <span class="tag-name">{{ item.tag }}</span>
        <span class="tag-count">{{ item.number }}</span>
      </span>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {type ArticleTag, articleTags} from "@/api/article";
import {PriceTag} from '@element-plus/icons-vue';

const tagTypes = ["primary", "success", "info", "warning", "danger"]

interface TagCloudItem {
  tag: string;
  number: number;
  type: string;
}

const tagCloudArray = ref<TagCloudItem[]>([])

const getTagCloudArray = async () => {
  let tagsArray: ArticleTag[]
  const res = await articleTags()
  if (res.code === 0) {
    tagsArray = res.data
    for (let i = 0; i < tagsArray.length; i++) {
      const item = tagsArray[i];
      const tagCloud: TagCloudItem = {
        tag: item.tag,
        number: item.number,
        type: tagTypes[i % tagTypes.length]
      }
      tagCloudArray.value.push(tagCloud);
    }
  }
}

getTagCloudArray()

const handleSearchJumps = (tag: string) => {
  window.open("/search?tag=" + tag)
}
</script>

<style scoped lang="scss">
.tag-cloud {
  background: var(--bg-card);
  border-radius: var(--radius-lg);
  border: none;
  box-shadow: var(--shadow-sm);
  overflow: hidden;
  
  :deep(.el-card__body) {
    padding: 0;
  }
  
  .section-header {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 16px 20px;
    background: linear-gradient(135deg, var(--color-accent) 0%, var(--color-accent-light) 100%);
    color: #fff;
    
    .header-icon {
      font-size: 1.25rem;
    }
    
    .title {
      font-size: 1.1rem;
      font-weight: 600;
    }
  }
  
  .tag-container {
    padding: 20px;
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
    
    .tag-item {
      display: inline-flex;
      align-items: center;
      gap: 6px;
      padding: 6px 14px;
      border-radius: 20px;
      font-size: 0.85rem;
      cursor: pointer;
      transition: all 0.3s ease;
      
      &:hover {
        transform: translateY(-2px);
        box-shadow: var(--shadow-sm);
      }
      
      .tag-name {
        font-weight: 500;
      }
      
      .tag-count {
        padding: 2px 8px;
        border-radius: 10px;
        font-size: 0.75rem;
        background: rgba(255, 255, 255, 0.3);
      }
      
      &.tag-primary {
        background: rgba(90, 159, 143, 0.15);
        color: var(--color-primary);
        
        &:hover {
          background: var(--color-primary);
          color: #fff;
        }
        
        .tag-count {
          background: rgba(90, 159, 143, 0.2);
        }
      }
      
      &.tag-success {
        background: rgba(103, 194, 58, 0.15);
        color: #67c23a;
        
        &:hover {
          background: #67c23a;
          color: #fff;
        }
        
        .tag-count {
          background: rgba(103, 194, 58, 0.2);
        }
      }
      
      &.tag-info {
        background: rgba(144, 147, 153, 0.15);
        color: #909399;
        
        &:hover {
          background: #909399;
          color: #fff;
        }
        
        .tag-count {
          background: rgba(144, 147, 153, 0.2);
        }
      }
      
      &.tag-warning {
        background: rgba(230, 162, 60, 0.15);
        color: #e6a23c;
        
        &:hover {
          background: #e6a23c;
          color: #fff;
        }
        
        .tag-count {
          background: rgba(230, 162, 60, 0.2);
        }
      }
      
      &.tag-danger {
        background: rgba(245, 108, 108, 0.15);
        color: #f56c6c;
        
        &:hover {
          background: #f56c6c;
          color: #fff;
        }
        
        .tag-count {
          background: rgba(245, 108, 108, 0.2);
        }
      }
    }
  }
}
</style>
