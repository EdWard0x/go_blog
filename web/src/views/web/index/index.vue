<template>
  <div class="index">
    <web-navbar/>
    <carousel/>
    <el-container class="main-content">
      <div class="container">
        <el-main>
          <!-- <advertisement/> -->
          <!-- <daily-news/> -->
          <article-list/>
        </el-main>

        <el-aside>
          <calendar/>
          <profile-card/>
          <tag-cloud/>
          <recent-comments/>
          <!-- <feedback/> -->
        </el-aside>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import Carousel from "@/components/layout/Carousel.vue";
import WebNavbar from "@/components/layout/WebNavbar.vue";
// import Advertisement from "@/components/pages/Advertisement.vue";
// import DailyNews from "@/components/pages/DailyNews.vue";
import ArticleList from "@/components/pages/ArticleList.vue";
import Calendar from "@/components/pages/Calendar.vue";
import ProfileCard from "@/components/pages/ProfileCard.vue";
import TagCloud from "@/components/pages/TagCloud.vue";
import RecentComments from "@/components/pages/RecentComments.vue";
import Feedback from "@/components/pages/Feedback.vue";
</script>

<style scoped lang="scss">
.index {
  min-height: 100vh;
  background-color: var(--bg-primary);

  .main-content {
    display: flex;
    justify-content: center;
    padding: 40px 20px;
    background-color: var(--bg-primary);

    .container {
      display: flex;
      max-width: 1400px;
      width: 100%;
      gap: 30px;

      .el-main {
        flex: 0 0 70%;
        max-width: 70%;
        padding: 0;
        overflow: visible;
      }

      .el-aside {
        flex: 0 0 30%;
        max-width: 30%;
        padding: 0;
        
        & > * {
          margin-bottom: 20px;
          
          &:last-child {
            margin-bottom: 0;
          }
        }
      }
    }
  }
}

@media (max-width: 1024px) {
  .index {
    .main-content {
      .container {
        flex-direction: column;
        
        .el-main,
        .el-aside {
          flex: none;
          max-width: 100%;
          width: 100%;
        }
      }
    }
  }
}
</style>
