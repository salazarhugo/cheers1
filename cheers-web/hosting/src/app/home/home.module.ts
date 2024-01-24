import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {HomeComponent} from "./features/home/home.component";
import {MaterialModule} from "../material/material.module";
import {MatCardModule} from "@angular/material/card";
import { RouterModule } from '@angular/router';
import { TopbarComponent } from './ui/topbar/topbar.component';
import {TopbarModule} from "./ui/topbar/topbar.module";
import {PostItemModule} from "../posts/ui/post-item/post-item.module";


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
        PostItemModule,
        MatCardModule,
        TopbarModule,
    ],
    exports: [
        ...components,
    ],
})
export class HomeModule {
}
