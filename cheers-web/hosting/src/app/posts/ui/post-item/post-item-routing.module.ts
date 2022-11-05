import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {PostItemComponent} from "./post-item.component";

const routes: Routes = [{path: '', component: PostItemComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class PostItemRoutingModule { }
