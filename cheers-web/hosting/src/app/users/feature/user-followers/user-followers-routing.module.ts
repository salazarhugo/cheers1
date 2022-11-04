import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {UserFollowersComponent} from "./user-followers.component";

const routes: Routes = [{path: '', component: UserFollowersComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class UserFollowersRoutingModule { }
