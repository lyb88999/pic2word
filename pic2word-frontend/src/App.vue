<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <el-container class="app-container">
    <el-header>
      <div class="header-container">
        <div class="logo">
          <h1>Pic2Word</h1>
        </div>
        <el-menu
          mode="horizontal"
          router
          :ellipsis="false"
          class="nav-menu"
        >
          <el-menu-item index="/">首页</el-menu-item>
          <el-menu-item index="/guide">使用指南</el-menu-item>
          <el-menu-item index="/about">关于我们</el-menu-item>
        </el-menu>
      </div>
    </el-header>
    
    <el-main>
      <RouterView />
    </el-main>
    
    <el-footer>
      <div class="footer-content">
        <p>© 2024 Pic2Word. 保留所有权利。</p>
      </div>
    </el-footer>
  </el-container>
</template>

<style>
/* 全局样式 */
body {
  margin: 0;
  padding: 0;
  font-family: 'Helvetica Neue', Helvetica, 'PingFang SC', 'Hiragino Sans GB', 'Microsoft YaHei', Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}

/* 组件样式 */
.app-container {
  min-height: 100vh;
}

.header-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
}

.logo h1 {
  margin: 0;
  color: #409eff;
  font-size: 24px;
}

.nav-menu {
  border-bottom: none;
}

.el-header {
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  background-color: white;
  position: sticky;
  top: 0;
  z-index: 1000;
}

.el-main {
  padding: 20px;
  background-color: #f5f7fa;
}

.el-footer {
  background-color: #f5f7fa;
  border-top: 1px solid #e6e6e6;
  padding: 20px;
  text-align: center;
  color: #909399;
}

.footer-content {
  max-width: 1200px;
  margin: 0 auto;
}
</style>
