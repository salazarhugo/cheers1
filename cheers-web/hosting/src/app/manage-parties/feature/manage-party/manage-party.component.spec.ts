import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartyComponent } from './manage-party.component';

describe('ManagePartyComponent', () => {
  let component: ManagePartyComponent;
  let fixture: ComponentFixture<ManagePartyComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartyComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartyComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
