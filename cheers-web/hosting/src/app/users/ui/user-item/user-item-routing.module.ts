import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UserItemComponent} from "./user-item.component";

const routes: Routes = [{path: '', component: UserItemComponent}];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class UserItemRoutingModule {
}
