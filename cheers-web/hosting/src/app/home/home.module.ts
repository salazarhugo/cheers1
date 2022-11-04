import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {HomeComponent} from "./features/home/home.component";
import {MaterialModule} from "../material/material.module";
import {PostsModule} from "../posts/posts.module";
import {MatCardModule} from "@angular/material/card";
import { RouterModule } from '@angular/router';
import { TopbarComponent } from './ui/topbar/topbar.component';
import {TopbarModule} from "./ui/topbar/topbar.module";


const components = [
    HomeComponent,
]

@NgModule({
    declarations: [
        HomeComponent,
        ...components,
    ],
    imports: [
        CommonModule,
        MaterialModule,
        RouterModule,
        PostsModule,
        MatCardModule,
        TopbarModule,
    ],
    exports: [
        ...components,
    ],
})
export class HomeModule {
}
