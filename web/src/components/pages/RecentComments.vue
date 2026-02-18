<template>
  <el-card class="recent-comments">
    <div class="section-header">
      <el-icon class="header-icon"><ChatDotRound /></el-icon>
      <span class="title">最新评论</span>
    </div>
    <div class="comments-container">
      <comment-item :comments="comments"/>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import CommentItem from "@/components/common/CommentItem.vue";
import {ref, watch} from "vue";
import {type Comment, commentNew} from "@/api/comment";
import {useLayoutStore} from "@/stores/layout";
import {ChatDotRound} from '@element-plus/icons-vue';

const comments=ref<Comment[]>([])

const getRecentCommentInfo = async ()=>{
  const res = await commentNew()
  if (res.code===0){
    comments.value=res.data
  }
}

getRecentCommentInfo()

const layoutStore = useLayoutStore()

watch(() => layoutStore.state.shouldRefreshCommentList, (newVal) => {
  if (newVal) {
    getRecentCommentInfo()
  }
})
</script>

<style scoped lang="scss">
.recent-comments {
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
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    color: #fff;
    
    .header-icon {
      font-size: 1.25rem;
    }
    
    .title {
      font-size: 1.1rem;
      font-weight: 600;
    }
  }
  
  .comments-container {
    padding: 16px;
    max-height: 400px;
    overflow-y: auto;
    
    &::-webkit-scrollbar {
      width: 4px;
    }
    
    &::-webkit-scrollbar-thumb {
      background: var(--border-color);
      border-radius: 2px;
    }
  }
}
</style>
