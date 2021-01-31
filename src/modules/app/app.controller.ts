import { Controller, Get } from '@nestjs/common';
import { AppService } from './app.service';
import { UserInfo, UserSkills } from '../../interfaces';

@Controller('/user')
export class AppController {
  constructor(private appService: AppService) {}

  @Get('/info')
  getUserInfo(): UserInfo {
    return this.appService.getUser();
  }

  @Get('/skills')
  getUserSkills(): UserSkills {
    return this.appService.getSkills();
  }
}
