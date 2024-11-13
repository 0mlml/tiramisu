#!/usr/bin/env python3

import requests
import json
from datetime import datetime
from colorama import init, Fore, Style
import time

# Initialize colorama for cross-platform colored output
init()

class APITester:
    def __init__(self):
        self.base_url = "http://localhost:8080/api"
        self.admin_token = None
        self.user_token = None
        self.admin_email = "admin@example.com"
        self.user_email = "user@example.com"
        self.question_ids = []
        self.submission_ids = []

    def print_result(self, success, message, error_response=None):
        """Helper function to print success/failure messages"""
        timestamp = datetime.now().strftime("%H:%M:%S")
        if success:
            print(f"{Fore.GREEN}[{timestamp}] ✓ {message}{Style.RESET_ALL}")
        else:
            print(f"{Fore.RED}[{timestamp}] ✗ {message}{Style.RESET_ALL}")
            if error_response:
                print(f"Error response: {error_response}")

    def register_admin(self):
        """Register an admin user"""
        print("\nRegistering admin user...")
        try:
            response = requests.post(
                f"{self.base_url}/register",
                json={
                    "email": self.admin_email,
                    "password": "admin123!@#",
                    "name": "Admin User"
                }
            )
            response.raise_for_status()
            data = response.json()
            self.admin_token = data['data']['token']
            self.print_result(True, "Admin registration successful")
        except Exception as e:
            if response.status_code == 400:
                self.print_result(True, "Admin user already exists")
            else:
                self.print_result(False, "Admin registration failed", str(e))

    def register_normal_user(self):
        """Register a normal user"""
        print("\nRegistering normal user...")
        try:
            response = requests.post(
                f"{self.base_url}/register",
                json={
                    "email": self.user_email,
                    "password": "user123!@#",
                    "name": "Normal User"
                }
            )
            response.raise_for_status()
            data = response.json()
            self.user_token = data['data']['token']
            self.print_result(True, "User registration successful")
        except Exception as e:
            if response.status_code == 400:
                self.print_result(True, "Normal user already exists")
            else:
                self.print_result(False, "User registration failed", str(e))

    def test_login_both_users(self):
        """Test login functionality for both users"""
        print("\nTesting login for both users...")
        
        # Test admin login
        try:
            response = requests.post(
                f"{self.base_url}/login",
                json={
                    "email": self.admin_email,
                    "password": "admin123!@#"
                }
            )
            response.raise_for_status()
            data = response.json()
            self.admin_token = data['data']['token']
            self.print_result(True, "Admin login successful")
        except Exception as e:
            self.print_result(False, "Admin login failed", str(e))

        # Test normal user login
        try:
            response = requests.post(
                f"{self.base_url}/login",
                json={
                    "email": self.user_email,
                    "password": "user123!@#"
                }
            )
            response.raise_for_status()
            data = response.json()
            self.user_token = data['data']['token']
            self.print_result(True, "User login successful")
        except Exception as e:
            self.print_result(False, "User login failed", str(e))

    def test_profile_operations(self):
        """Test profile related operations"""
        print("\nTesting profile operations...")
        
        # Test getting admin profile
        try:
            response = requests.get(
                f"{self.base_url}/profile",
                headers={"Authorization": f"Bearer {self.admin_token}"}
            )
            response.raise_for_status()
            self.print_result(True, "Get admin profile successful")
        except Exception as e:
            self.print_result(False, "Get admin profile failed", str(e))

        # Test updating admin profile
        try:
            response = requests.put(
                f"{self.base_url}/profile",
                headers={"Authorization": f"Bearer {self.admin_token}"},
                json={
                    "name": "Updated Admin Name",
                    "picture": "https://example.com/admin.jpg"
                }
            )
            response.raise_for_status()
            self.print_result(True, "Update admin profile successful")
        except Exception as e:
            self.print_result(False, "Update admin profile failed", str(e))

    def test_question_operations(self):
        """Test question related operations"""
        print("\nTesting question operations...")
        
        # Test creating questions (admin only)
        questions = [
            {
                "question": "How satisfied are you with our service?",
                "type": "scale",
                "min": 1,
                "max": 5
            },
            {
                "question": "Do you feel burned out?",
                "type": "scale",
                "min": 1,
                "max": 10
            }
        ]

        for q in questions:
            try:
                response = requests.post(
                    f"{self.base_url}/admin/questions",
                    headers={"Authorization": f"Bearer {self.admin_token}"},
                    json=q
                )
                response.raise_for_status()
                data = response.json()
                self.question_ids.append(data['data']['id'])
                self.print_result(True, f"Create question successful: {q['question']}")
            except Exception as e:
                self.print_result(False, f"Create question failed: {q['question']}", str(e))

        # Test getting questions (both users)
        for token in [self.admin_token, self.user_token]:
            try:
                response = requests.get(
                    f"{self.base_url}/questions",
                    headers={"Authorization": f"Bearer {token}"}
                )
                response.raise_for_status()
                self.print_result(True, "Get questions successful")
            except Exception as e:
                self.print_result(False, "Get questions failed", str(e))

    def test_submission_operations(self):
        """Test submission related operations"""
        print("\nTesting submission operations...")
        
        # Submit questionnaire as normal user
        try:
            answers = [
                {"id": self.question_ids[0], "answer": "4"},
                {"id": self.question_ids[1], "answer": "2"},
            ]
            
            response = requests.post(
                f"{self.base_url}/submit",
                headers={"Authorization": f"Bearer {self.user_token}"},
                json={"answers": answers}
            )
            response.raise_for_status()
            data = response.json()
            self.submission_ids.append(data['data']['id'])
            self.print_result(True, "Submit questionnaire successful")
        except Exception as e:
            self.print_result(False, "Submit questionnaire failed", str(e))

        # Test getting user submissions
        try:
            response = requests.get(
                f"{self.base_url}/submissions",
                headers={"Authorization": f"Bearer {self.user_token}"}
            )
            response.raise_for_status()
            self.print_result(True, "Get user submissions successful")
        except Exception as e:
            self.print_result(False, "Get user submissions failed", str(e))

        # Test getting all submissions (admin only)
        try:
            response = requests.get(
                f"{self.base_url}/admin/submissions/all",
                headers={"Authorization": f"Bearer {self.admin_token}"}
            )
            response.raise_for_status()
            self.print_result(True, "Get all submissions successful")
        except Exception as e:
            self.print_result(False, "Get all submissions failed", str(e))

    def test_unauthorized_access(self):
        """Test unauthorized access attempts"""
        print("\nTesting unauthorized access...")
        
        # Try to access admin endpoint with normal user token
        try:
            response = requests.get(
                f"{self.base_url}/admin/submissions/all",
                headers={"Authorization": f"Bearer {self.user_token}"}
            )
            success = response.status_code == 403  # Expected to fail with forbidden
            self.print_result(success, "Unauthorized access test successful" if success else "Unauthorized access test failed")
        except Exception as e:
            self.print_result(True, "Unauthorized access properly denied")

    def cleanup(self):
        """Clean up created questions"""
        print("\nCleaning up...")
        
        # Delete test questions
        for qid in self.question_ids:
            try:
                response = requests.delete(
                    f"{self.base_url}/admin/questions/{qid}",
                    headers={"Authorization": f"Bearer {self.admin_token}"}
                )
                response.raise_for_status()
                self.print_result(True, f"Deleted question {qid}")
            except Exception as e:
                self.print_result(False, f"Failed to delete question {qid}", str(e))

    def run_all_tests(self):
        """Run all tests in sequence"""
        test_sequence = [
            self.register_admin,
            self.register_normal_user,
            self.test_login_both_users,
            self.test_profile_operations,
            self.test_question_operations,
            self.test_submission_operations,
            self.test_unauthorized_access,
            self.cleanup
        ]

        for test in test_sequence:
            test()
            time.sleep(0.5)  # Small delay between tests to avoid rate limiting

def main():
    print(f"{Fore.CYAN}Starting API tests...{Style.RESET_ALL}")
    tester = APITester()
    tester.run_all_tests()
    print(f"\n{Fore.CYAN}Tests completed!{Style.RESET_ALL}")

if __name__ == "__main__":
    main()