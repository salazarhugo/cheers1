import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {PostComponent} from "./features/post/post.component";
import {MaterialModule} from "../material/material.module";
import {RouterModule} from "@angular/router";


const components = [
    PostComponent,
]

@NgModule({
    declarations: [
        ...components,
    ],
    imports: [
        CommonModule,
        MaterialModule,
        RouterModule,
    ],
    exports: [
        ...components,
    ]
})
export class PostsModule {
}
