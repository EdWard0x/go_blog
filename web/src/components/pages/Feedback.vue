<template>
  <el-card class="feedback">
    <div class="section-header">
      <el-icon class="header-icon"><EditPen /></el-icon>
      <span class="title">意见反馈</span>
    </div>
    <div class="feedback-content">
      <el-input 
        type="textarea" 
        :rows="3" 
        v-model="feedbackCreateFormData.content" 
        maxlength="100" 
        show-word-limit
        placeholder="请输入反馈建议..."
        resize="none"
      />
      <div class="feedback-actions">
        <el-text class="tip">
          <el-icon><InfoFilled /></el-icon>
          请登录后再进行反馈
        </el-text>
        <div class="button-group">
          <el-button @click="feedbackCreateFormData.content=''" size="default">取消</el-button>
          <el-button @click="submitForm" type="primary" size="default">提交</el-button>
        </div>
      </div>
    </div>
    
    <div class="feedback-list" v-if="feedbackInfoList.length > 0">
      <div class="list-header">
        <span>最新反馈</span>
      </div>
      <div class="feedback-items">
        <div class="feedback-item" v-for="item in feedbackInfoList" :key="item.time">
          <div class="item-content">{{ item.content }}</div>
          <div class="item-meta">
            <span class="item-time">{{ item.time }}</span>
          </div>
          <div class="item-reply" v-if="item.reply">
            <el-icon><ChatLineSquare /></el-icon>
            <span>回复：{{ item.reply }}</span>
          </div>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import {reactive, ref, watch} from "vue";
import {feedbackCreate, type FeedbackCreateRequest, feedbackNew} from "@/api/feedback";
import {EditPen, InfoFilled, ChatLineSquare} from '@element-plus/icons-vue';

const feedbackCreateFormData = reactive<FeedbackCreateRequest>({
  content: '',
})

interface FeedbackNew {
  content: string;
  reply: string;
  time: string;
}

const feedbackInfoList = ref<FeedbackNew[]>([])

const shouldRefreshFeedbackInfoTable = ref(false)
watch(() => shouldRefreshFeedbackInfoTable.value, (newVal) => {
  if (newVal) {
    getFeedbackNew()
    shouldRefreshFeedbackInfoTable.value = false;
  }
})

const submitForm = async () => {
  const res = await feedbackCreate(feedbackCreateFormData)
  if (res.code === 0) {
    ElMessage.success(res.msg)
    feedbackCreateFormData.content = ''
    shouldRefreshFeedbackInfoTable.value = true
  }
}

const getFeedbackNew = async () => {
  feedbackInfoList.value = []
  const res = await feedbackNew()
  if (res.code === 0) {
    res.data.forEach(value => {
      const date = new Date(value.created_at);
      const info: FeedbackNew = {
        content: value.content,
        reply: value.reply,
        time: date.toLocaleString(),
      }
      feedbackInfoList.value.push(info)
    })
  }
}

getFeedbackNew()
</script>

<style scoped lang="scss">
.feedback {
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
    background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
    color: #fff;
    
    .header-icon {
      font-size: 1.25rem;
    }
    
    .title {
      font-size: 1.1rem;
      font-weight: 600;
    }
  }
  
  .feedback-content {
    padding: 20px;
    
    .el-textarea {
      :deep(.el-textarea__inner) {
        background: var(--bg-secondary);
        border-color: var(--border-color);
        color: var(--text-primary);
        border-radius: var(--radius-md);
        
        &:focus {
          border-color: var(--color-primary);
        }
      }
    }
    
    .feedback-actions {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-top: 12px;
      
      .tip {
        display: flex;
        align-items: center;
        gap: 4px;
        color: var(--text-muted);
        font-size: 0.85rem;
      }
      
      .button-group {
        display: flex;
        gap: 8px;
      }
    }
  }
  
  .feedback-list {
    border-top: 1px solid var(--border-color-light);
    
    .list-header {
      padding: 12px 20px;
      font-weight: 500;
      color: var(--text-secondary);
      font-size: 0.9rem;
      background: var(--bg-secondary);
    }
    
    .feedback-items {
      max-height: 300px;
      overflow-y: auto;
      
      &::-webkit-scrollbar {
        width: 4px;
      }
      
      &::-webkit-scrollbar-thumb {
        background: var(--border-color);
        border-radius: 2px;
      }
      
      .feedback-item {
        padding: 16px 20px;
        border-bottom: 1px solid var(--border-color-light);
        transition: background 0.2s ease;
        
        &:last-child {
          border-bottom: none;
        }
        
        &:hover {
          background: var(--bg-secondary);
        }
        
        .item-content {
          color: var(--text-primary);
          font-size: 0.9rem;
          line-height: 1.5;
          margin-bottom: 8px;
        }
        
        .item-meta {
          display: flex;
          justify-content: flex-end;
          
          .item-time {
            font-size: 0.8rem;
            color: var(--text-muted);
          }
        }
        
        .item-reply {
          display: flex;
          align-items: flex-start;
          gap: 6px;
          margin-top: 10px;
          padding: 10px 12px;
          background: rgba(90, 159, 143, 0.1);
          border-radius: var(--radius-sm);
          font-size: 0.85rem;
          color: var(--text-secondary);
          
          .el-icon {
            color: var(--color-primary);
            margin-top: 2px;
          }
        }
      }
    }
  }
}
</style>
