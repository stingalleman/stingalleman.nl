import { Injectable } from '@nestjs/common';
import { UserInfo, UserSkills } from '../../interfaces';
import config from '../../util/user.json';

@Injectable()
export class AppService {
  getUser(): UserInfo {
    return config.user;
  }

  getSkills(): UserSkills {
    return { skills: config.skills };
  }
}
