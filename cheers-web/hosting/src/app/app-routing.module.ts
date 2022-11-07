import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {UserResolver} from "./users/data/user.resolver";
import {CoreComponent} from "./core/feature/core.component";
import {SignInComponent} from "./shared/ui/sign-in/sign-in.component";
import {FinishSignUpComponent} from "./shared/ui/finish-sign-up/finish-sign-up.component";
import {TicketingComponent} from "./shared/ui/ticketing/ticketing.component";
import {PaymentComponent} from "./payments/features/payment/payment.component";
import {CompleteComponent} from "./shared/ui/complete/complete.component";
import {StoriesComponent} from "./stories/features/story-feed/stories.component";
import {ManagePartyComponent} from "./manage-parties/feature/manage-party/manage-party.component";
import {PartyResolver} from "./manage-parties/data/party.resolver";

const routes: Routes = [
    {
        path: 'finish-sign-up',
        component: FinishSignUpComponent,
    },
    {
        path: 'payment/complete',
        component: CompleteComponent,
    },
    {
        path: 'payment/:clientSecret',
        component: PaymentComponent,
    },
    {
        path: 'story-feed/:id',
        component: StoriesComponent,
    },
    {
        path: 'ticketing/:id',
        component: TicketingComponent,
    },
    {
        path: 'p/:id',
        loadChildren: () => import('./parties/feature/party-detail/party-detail.module').then(m => m.PartyDetailModule)
    },
    {
        path: 'sign-in',
        component: SignInComponent,
    },
    {
        path: 'manage/parties',
        component: ManagePartyComponent, resolve: {party: PartyResolver},
        loadChildren: () => import('./manage-parties/feature/manage-party-shell/manage-party-shell.module').then(m => m.ManagePartyShellModule)
    },
    {
        path: '',
        component: CoreComponent, resolve: {data: UserResolver},
        loadChildren: () => import('./core/core.module').then(m => m.CoreModule)
    },
    {path: '**', redirectTo: 'sign-in'},
];

@NgModule({
    declarations: [],
    imports: [RouterModule.forRoot(routes)],
    exports: [RouterModule]
})
export class AppRoutingModule {
}
