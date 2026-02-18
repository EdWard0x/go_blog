<template>
  <el-card class="calendar">
    <div class="calendar-header">
      <el-icon class="calendar-icon"><Calendar /></el-icon>
      <span class="title">今日日历</span>
    </div>
    <div class="calendar-content">
      <div class="time-display">
        <span class="current-time">{{ currentTime }}</span>
        <span class="current-date">{{ calendarInfo.date }}</span>
      </div>
      <div class="calendar-info">
        <div class="info-item">
          <span class="info-label">农历</span>
          <span class="info-value">{{ calendarInfo.lunar_date }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">干支</span>
          <span class="info-value">{{ calendarInfo.ganzhi }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">星座</span>
          <span class="info-value">{{ calendarInfo.zodiac }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">节气</span>
          <span class="info-value">{{ calendarInfo.solar_term }}</span>
        </div>
        <div class="info-item">
          <span class="info-label">天次</span>
          <span class="info-value">{{ calendarInfo.day_of_year }}</span>
        </div>
      </div>
      <div class="calendar-footer">
        <div class="fortune-item auspicious">
          <span class="fortune-label">宜</span>
          <span class="fortune-value">{{ calendarInfo.auspicious }}</span>
        </div>
        <div class="fortune-item inauspicious">
          <span class="fortune-label">忌</span>
          <span class="fortune-value">{{ calendarInfo.inauspicious }}</span>
        </div>
      </div>
    </div>
  </el-card>
</template>

<script setup lang="ts">
import {onUnmounted, ref} from "vue";
import {websiteCalendar, type WebsiteCalendarResponse} from "@/api/website";
import {Calendar} from '@element-plus/icons-vue';

const calendarInfo = ref<WebsiteCalendarResponse>({
  date: '',
  lunar_date: '',
  ganzhi: '',
  zodiac: '',
  day_of_year: '',
  solar_term: '',
  auspicious: '',
  inauspicious: '',
})

let timerId: number | null = null
const currentTime = ref('')

function updateCurrentTime() {
  const now = new Date()
  currentTime.value = now.toLocaleTimeString('zh-CN', { hour: '2-digit', minute: '2-digit', second: '2-digit' })
}

function initializeTimer() {
  updateCurrentTime()
  timerId = setInterval(updateCurrentTime, 1000)
}

onUnmounted(() => {
  clearInterval(timerId as number)
})

initializeTimer()

const getCalendarInfo = async () => {
  const res = await websiteCalendar()
  if (res.code == 0) {
    calendarInfo.value = res.data
  }
}
getCalendarInfo()
</script>

<style scoped lang="scss">
.calendar {
  background: var(--bg-card);
  border-radius: var(--radius-lg);
  border: none;
  box-shadow: var(--shadow-sm);
  overflow: hidden;
  
  :deep(.el-card__body) {
    padding: 0;
  }
  
  .calendar-header {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 16px 20px;
    background: linear-gradient(135deg, var(--color-primary) 0%, var(--color-primary-light) 100%);
    color: #fff;
    
    .calendar-icon {
      font-size: 1.25rem;
    }
    
    .title {
      font-size: 1.1rem;
      font-weight: 600;
    }
  }
  
  .calendar-content {
    padding: 20px;
    
    .time-display {
      text-align: center;
      padding: 16px;
      background: var(--bg-secondary);
      border-radius: var(--radius-md);
      margin-bottom: 16px;
      
      .current-time {
        display: block;
        font-size: 1.75rem;
        font-weight: 600;
        color: var(--color-primary);
        font-family: 'Courier New', monospace;
        letter-spacing: 2px;
      }
      
      .current-date {
        display: block;
        font-size: 0.9rem;
        color: var(--text-secondary);
        margin-top: 4px;
      }
    }
    
    .calendar-info {
      margin-bottom: 16px;
      
      .info-item {
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding: 10px 0;
        border-bottom: 1px dashed var(--border-color-light);
        
        &:last-child {
          border-bottom: none;
        }
        
        .info-label {
          color: var(--text-muted);
          font-size: 0.9rem;
        }
        
        .info-value {
          color: var(--text-primary);
          font-weight: 500;
        }
      }
    }
    
    .calendar-footer {
      display: flex;
      gap: 12px;
      
      .fortune-item {
        flex: 1;
        padding: 12px;
        border-radius: var(--radius-sm);
        text-align: center;
        
        .fortune-label {
          display: block;
          font-size: 0.85rem;
          font-weight: 600;
          margin-bottom: 4px;
        }
        
        .fortune-value {
          display: block;
          font-size: 0.8rem;
          line-height: 1.4;
        }
        
        &.auspicious {
          background: rgba(103, 194, 58, 0.1);
          
          .fortune-label {
            color: #67c23a;
          }
          
          .fortune-value {
            color: var(--text-secondary);
          }
        }
        
        &.inauspicious {
          background: rgba(245, 108, 108, 0.1);
          
          .fortune-label {
            color: #f56c6c;
          }
          
          .fortune-value {
            color: var(--text-secondary);
          }
        }
      }
    }
  }
}
</style>
