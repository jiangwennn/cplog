<template>
    <div class="flex h-full">
        <div class="basis-80 flex-none h-full flex select-none">
            <div class="w-full h-full bg-white box-border border-r border-neutral-200">
                <div class="w-full h-16 bg-white flex items-center box-border border-b border-neutral-100">
                    <div class="basis-5/6 h-8 ml-2 flex items-center relative">
                        <input type="text" placeholder="搜索历史记录" class="w-full h-full rounded-md indent-7 text-xs bg-zinc-100 focus:outline-none focus:ring-0 focus:border focus:border-[#df1a29]" />
                        <IconSearch class="absolute ml-2 top-[3px]" size="14" stroke-width="1" />
                    </div>
                    <div class="basis-1/6 text-center">
                        <IconPlus class="bg-zinc-100 block h-full w-[26px] mx-auto rounded-sm active:bg-zinc-50" theme="outline" size="24" fill="#333" stroke-width="2"/>
                    </div>
                </div>
                <div class="overflow-y-auto h-full">
                    <div v-for="item in listData" class="w-full h-24 pt-2 pb-2 border-b border-neutral-100 text-xs active:bg-neutral-50">
                        <div class="relative w-full h-11 text-neutral-600 p-3 box-border line-clamp-2 ">
                            {{ item.content }}
                        </div>
                        <div class="relative w-full h-4 flex p-3 pb-0">
                            <span class="basis-1/2 text-neutral-400 text-xs"> {{ item.time }}</span>
                            <icon class="basis-1/2 flex items-center">
                                <IconLike class="basis-1/2 flex-none" />
                            </icon>
                        </div>
                    </div>
                </div>
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
import {
  Plus as IconPlus,
  Search as IconSearch,
  Like as IconLike,
} from '@icon-park/vue-next';

let listData = [
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '首先，你需要给div添加一个overflow-y类，比如overflow-y-auto，来指定当div的内容超出高度时显示垂直滚动条。1',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '只需书写 HTML 代码，无需书写 CSS，即可快速构建美观的网站。',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '指定当div的内容超出高度时显示垂直滚动条。1',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '本CSS框架本质上是一个工具集，包含了大量类似 flex、 pt-4、 text-center 以及 rotate-90 等工具类，可以组合使用并直接在 HTML 代码上实现任何 UI 设计。',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '我曾经撰写了 数千字的文章 来解释为什么传统的“语义化命名”是 CSS 很难维护的原因， 但现实是，在你真正遇到问题之前你是永远不会相信我的。如果你 能忍住放弃的冲动来尝试这个框架的话，我相信 你会后悔自己为什么曾经用其他的方式写 CSS。',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '“最佳实践”实际上没有用。',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '“最佳实践”实际上没有用。',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '“最佳实践”实际上没有用。',
    },
    {
        id: '1',
        time: '2023-07-04 23:05:35',
        content: '“最佳实践”实际上没有用。',
    },
];


</script>

<style></style>