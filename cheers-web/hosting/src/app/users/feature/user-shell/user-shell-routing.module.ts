import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {OtherProfileComponent} from "../user-profile/other-profile.component";

const routes: Routes = [
    {
        path: 'edit',
        loadChildren: () => import('../user-edit/edit.module').then(m => m.EditModule),
    },
    {
        path: ':username/followers',
        loadChildren: () => import('../user-followers/user-followers.module').then(m => m.UserFollowersModule),
    },
];

@NgModule({
    imports: [RouterModule.forChild(routes)],
    exports: [RouterModule]
})
export class UserShellRoutingModule {
}
