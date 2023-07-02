<template>
    <div class="flex h-full">
        <div class="basis-80 flex-none h-full flex select-none">
            <div class="w-full h-full bg-white">
                我是列表页
            </div>
        </div>
        <!-- 右侧详情 -->
        <div class="h-full flex-auto bg-gray-100 pl-4" style="--wails-draggable:drag">
            <!-- 主页介绍 -->
            <About class="h-full mx-auto w-full select-none"></About>
        </div>
    </div>
</template>

<script setup>
import About from '@/components/About.vue';

</script>

<style></style>